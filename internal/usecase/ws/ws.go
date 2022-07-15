package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"web-socket/internal/entity"
	"web-socket/pkg/response"

	"github.com/gorilla/websocket"
)

func ServerWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		response.HandleMessage(w, http.StatusUpgradeRequired, err)
	}
	if has, clientId := r.URL.Query().Has("client-id"), r.URL.Query().Get("client-id"); has {
		fmt.Println("the client ID", clientId)
		hub.AddClient(clientId, &Client{clientId: clientId, conn: conn, hub: hub})
	} else {
		response.HandleMessage(w, http.StatusBadRequest, errors.New("client-id is required"))
	}
}

type Hub struct {
	mu      sync.Mutex
	clients map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

func (h *Hub) AddClient(clientId string, client *Client) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.clients[clientId] = client
	go client.Read()
}

func (h *Hub) SendMessage(payload entity.Message) error {
	defer h.mu.Unlock()
	h.mu.Lock()
	if client, ok := h.clients[payload.RecipientId]; ok {
		msg, _ := json.Marshal(payload)
		return client.Write(msg)
	}
	return fmt.Errorf("ws - client not found")
}

type Client struct {
	conn     *websocket.Conn
	hub      *Hub
	clientId string
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) Read() {
	msgChan := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Print("err: ", err)
				}
				log.Print("not covered err: ", err)
				if err := c.Close(); err != nil {
					log.Print("close err: ", err)
				}
				return
			}
			msgChan <- strings.ToLower(string(bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))))
		}
	}(ctx)
	for {
		select {
		case msg := <-msgChan:
			if msg == "ping" {
				c.Write([]byte("pong"))
			}
		case <-time.After(time.Second * 3):
			if err := c.Close(); err != nil {
				log.Print("close err: ", err)
			}
			cancel()
			return
		}
	}
}

func (c *Client) Close() error {
	delete(c.hub.clients, c.clientId)
	return c.conn.Close()
}

func (c *Client) Write(message []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, message)

}
