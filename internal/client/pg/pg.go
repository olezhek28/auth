package pg

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	Name     string
	QueryRaw string
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type NamedExecer interface {
	GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type PG interface {
	QueryExecer
	NamedExecer
	Pinger
	Close() error
}

type pg struct {
	pgxPool *pgxpool.Pool
}

func (p *pg) GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pgxPool.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	return p.pgxPool.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	return p.pgxPool.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) Close() error {
	p.pgxPool.Close()
	return nil
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pgxPool.Ping(ctx)
}
