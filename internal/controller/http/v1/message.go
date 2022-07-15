package v1

import (
	"net/http"
	"web-socket/internal/entity"
	"web-socket/internal/usecase"
	"web-socket/internal/usecase/ws"
	"web-socket/pkg/logger"

	"github.com/gin-gonic/gin"
)

type messageRoutes struct {
	muc *usecase.MessageUseCase
	l   logger.Interface
}

func newMessageRoutes(handler *gin.RouterGroup, uc *usecase.MessageUseCase, l logger.Interface) {
	m := &messageRoutes{uc, l}
	h := handler.Group("/message")
	{
		h.POST("/send", m.sendMessage)
		h.GET("/ws", func(c *gin.Context) {
			ws.ServerWS(uc.GetHub(), c.Writer, c.Request)
		})
	}

}

type sendMessageRequest struct {
	SenderId    string `json:"sender_id"`
	RecipientId string `json:"recipient_id"`
	Message     string `json:"message"`
}

// @Summary		Send Messsage
// @Description	Send Message to Client
// @ID			send-message
// @Tags		message
// @Accept		json
// @Produce		json
// @Param		message body sendMessageRequest true "Message"
// @Success		200 {object} entity.Message
// @Failure		500 {object} response
// @Router		/v1/message/send [post]
func (m *messageRoutes) sendMessage(c *gin.Context) {
	var request sendMessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		m.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
	message, err := m.muc.SendMessage(c.Request.Context(), entity.Message{
		SenderId:    request.SenderId,
		RecipientId: request.RecipientId,
		Message:     request.Message,
	})
	if err != nil {
		m.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "failed to save message")

		return
	}

	c.JSON(http.StatusCreated, message)
}
