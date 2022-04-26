package model

type User struct {
	UserID   string `db:"user_id"`
	Name     string `db:"name"`
	Birthday string `db:"birthday"`
}
