package usecase

import (
	"context"
	"web-socket/internal/entity"
)

type (
	// MessageRepository is an interface for message repository
	MessageRepo interface {
		Create(context.Context, entity.Message) (entity.Message, error)
	}
)
