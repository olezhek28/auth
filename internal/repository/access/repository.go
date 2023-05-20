package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/auth/internal/client/pg"
	"github.com/olezhek28/auth/internal/model"
)

type Repository interface {
	GetList(ctx context.Context) ([]*model.AccessInfo, error)
}

type repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) GetList(ctx context.Context) ([]*model.AccessInfo, error) {
	builder := sq.Select("id", "endpoint_address", "role", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From("accesses")

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "accesses.GetList",
		QueryRaw: query,
	}

	var accessInfo []*model.AccessInfo
	err = r.client.PG().SelectContext(ctx, accessInfo, q, v...)
	if err != nil {
		return nil, err
	}

	return accessInfo, nil
}
