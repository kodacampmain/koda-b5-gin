package model

type User struct {
	Id     int    `db:"id"`
	Gender string `db:"gender"`
	Name   string `db:"name"`
}
