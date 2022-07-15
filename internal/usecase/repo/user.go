package repo

import (
	"context"
	"web-socket/internal/entity"
	"web-socket/pkg/postgres"

	sq "github.com/Masterminds/squirrel"
)

type userRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *userRepo {
	return &userRepo{pg}
}

func (r *userRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	sql, params, err := r.Builder.Insert("user").
		Columns("name", "username", "password").
		Values(user.Name, user.Username, user.Password).Suffix("RETURNING *").ToSql()
	if err != nil {
		return entity.User{}, err
	}
	var result entity.User
	err = r.Pool.QueryRow(ctx, sql, params...).Scan(&result.Id, &result.Name, &result.Username, &result.Password)
	return entity.User{}, err
}

func (r *userRepo) Get(ctx context.Context, user entity.User) (entity.User, error) {
	sql, params, err := r.Builder.Select("").From("user").Where(sq.Eq{"id": user.Id}).ToSql()
	if err != nil {
		return entity.User{}, err
	}
	var result entity.User
	err = r.Pool.QueryRow(ctx, sql, params...).Scan(&result.Id, &result.Name, &result.Username, &result.Password)
	return result, err
}
