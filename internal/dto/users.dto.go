package dto

type User struct {
	Id     int    `json:"id"`
	Gender string `json:"jenis_kelamin"`
	Name   string `json:"name"`
}

type NewUser struct {
	Name   string `json:"name" binding:"required"`
	Gender string `json:"gender" binding:"required"`
}
