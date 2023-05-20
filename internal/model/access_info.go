package model

import (
	"database/sql"
	"time"
)

type AccessInfo struct {
	ID              int64        `db:"id"`
	EndpointAddress string       `db:"endpoint_address"`
	Role            string       `db:"role"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}
