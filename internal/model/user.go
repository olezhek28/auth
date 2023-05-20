package model

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
