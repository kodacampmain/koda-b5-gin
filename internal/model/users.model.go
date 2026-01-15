package model

type User struct {
	Id         int     `db:"id"`
	Gender     string  `db:"gender"`
	Email      string  `db:"email"`
	ProfileImg *string `db:"profile_img"`
	Password   string  `db:"password"`
	Role       string  `db:"role"`
}
