package model

import (
	"database/sql"
	"time"
)

type Info struct {
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type Note struct {
	ID        int64        `db:"id"`
	Info      Info         `db:""`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
