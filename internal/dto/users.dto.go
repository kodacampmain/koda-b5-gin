package dto

import "mime/multipart"

type User struct {
	Id     int    `json:"id"`
	Gender string `json:"jenis_kelamin"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type NewUser struct {
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EditUser struct {
	Image *multipart.FileHeader `form:"image"`
}

type EditPassword struct {
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
