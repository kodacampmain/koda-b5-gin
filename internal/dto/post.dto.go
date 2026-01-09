package dto

type PostBody struct {
	Name string `form:"name" example:"fadhlul"`
	Age  int    `form:"age" example:"24"`
}
