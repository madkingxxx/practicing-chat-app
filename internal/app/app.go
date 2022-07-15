package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"web-socket/config"
	v1 "web-socket/internal/controller/http/v1"
	"web-socket/internal/usecase"
	"web-socket/internal/usecase/repo"
	"web-socket/internal/usecase/ws"
	"web-socket/pkg/httpserver"
	"web-socket/pkg/logger"
	"web-socket/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.Load()
	l := logger.New("info")

	pg, err := postgres.New(cfg.Postgres.URL)
	if err != nil {
		l.Fatal(err)
	}

	messageUseCase := usecase.New(
		repo.NewMessageRepo(pg),
		ws.NewHub(),
	)

	handler := gin.New()
	v1.NewRouter(handler, messageUseCase, l)
	httpServer := httpserver.New(handler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: ", s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
