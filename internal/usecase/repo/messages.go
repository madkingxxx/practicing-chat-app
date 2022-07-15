package repo

import (
	"context"
	"web-socket/internal/entity"
	"web-socket/pkg/postgres"
)

type messageRepo struct {
	*postgres.Postgres
}

func NewMessageRepo(pg *postgres.Postgres) *messageRepo {
	return &messageRepo{pg}
}

func (r *messageRepo) Create(ctx context.Context, message entity.Message) (entity.Message, error) {
	sql, params, err := r.Builder.Insert("messages").
		Columns("sender_id", "recipient_id", "message").
		Values(message.SenderId, message.RecipientId, message.Message).ToSql()
	if err != nil {
		return entity.Message{}, err
	}
	_, err = r.Pool.Exec(ctx, sql, params...)
	return entity.Message{}, err
}
