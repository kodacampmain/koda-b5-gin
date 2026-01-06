package dto

type PostBody struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}
