package note

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olezhek28/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "note"

type Repository interface {
	Create(ctx context.Context, info *model.Info) (int64, error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Create(ctx context.Context, info *model.Info) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "content", "created_at").
		Values(info.Title, info.Content, info.CreatedAt).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := r.pool.Query(ctx, query, v...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	rows.Next()
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
