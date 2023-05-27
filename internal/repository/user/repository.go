package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/auth/internal/client/pg"
	"github.com/olezhek28/auth/internal/model"
)

type Repository interface {
	Get(ctx context.Context, username string) (*model.User, error)
}

type repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Get(ctx context.Context, username string) (*model.User, error) {
	builder := sq.Select("username", "password", "role").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"username": username}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "users.Get",
		QueryRaw: query,
	}

	user := new(model.User)
	err = r.client.PG().GetContext(ctx, user, q, v...)
	if err != nil {
		return nil, err
	}

	return user, nil
}
