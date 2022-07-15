package usecase

import (
	"context"
	"fmt"
	"web-socket/internal/entity"
	"web-socket/internal/usecase/ws"
)

type MessageUseCase struct {
	repo MessageRepo
	hub  *ws.Hub
}

// New Message Use Case ...
func New(r MessageRepo, hub *ws.Hub) *MessageUseCase {
	return &MessageUseCase{
		repo: r,
		hub:  hub,
	}
}

// Send Message ...
func (uc *MessageUseCase) SendMessage(ctx context.Context, msg entity.Message) (entity.Message, error) {
	err := uc.hub.SendMessage(msg)
	if err != nil {
		return entity.Message{}, err
	}

	message, err := uc.repo.Create(ctx, msg)
	if err != nil {
		return entity.Message{}, fmt.Errorf("MessageUseCase - SendMessage - uc.repo.Create: %w", err)
	}

	return message, nil
}

func (uc *MessageUseCase) GetHub() *ws.Hub {
	return uc.hub
}
