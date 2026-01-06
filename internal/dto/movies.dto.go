package dto

// property datatype struct_tag

type MoviesParam struct {
	Id   int    `uri:"id"`
	Slug string `uri:"slug"`
}

type MoviesQuery struct {
	Title string   `form:"title"`
	Genre []string `form:"genre"`
	Page  int      `form:"page"`
}
