package note

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/auth/internal/client/pg"
	"github.com/olezhek28/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "note"

type Repository interface {
	Create(ctx context.Context, info *model.Info) (int64, error)
	GetList(ctx context.Context) ([]*model.Note, error)
}

type repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *repository {
	return &repository{
		client: client,
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

	q := pg.Query{
		Name:     "note.Create",
		QueryRaw: query,
	}

	rows, err := r.client.PG().QueryContext(ctx, q, v...)
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

func (r *repository) GetList(ctx context.Context) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "content", "created_at", "updated_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "note.GetList",
		QueryRaw: query,
	}

	var notes []*model.Note
	err = r.client.PG().SelectContext(ctx, &notes, q, v...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
