package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialization
	app := gin.Default()
	// routing
	app.GET("/", func(c *gin.Context) {
		whoami := c.GetHeader("x-who-am-i")
		c.JSON(http.StatusOK, gin.H{
			"msg":    "Selamat Datang",
			"whoami": whoami,
		})
	})
	app.POST("/", func(c *gin.Context) {
		// reading body
		var body PostBody
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
				// "err": err.Error(),
			})
			return
		}
		// validasi
		c.JSON(http.StatusOK, gin.H{
			"msg":  "OK",
			"data": body,
		})
	})
	app.GET("/movies/:id/:slug", func(c *gin.Context) {
		// direct access
		id, _ := strconv.Atoi(c.Param("id"))
		slug := c.Param("slug")
		// data binding
		var moviesParam MoviesParam
		if err := c.ShouldBindUri(&moviesParam); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
			"params": gin.H{
				"id":   id,
				"slug": slug,
			},
			"moviesParam": moviesParam,
		})
	})
	app.GET("/movies", func(c *gin.Context) {
		// direct access
		title := c.Query("title")
		genre := c.QueryArray("genre")
		page := c.Query("page")
		// data binding
		var moviesQuery MoviesQuery
		if err := c.ShouldBindQuery(&moviesQuery); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
			"query": gin.H{
				"title": title,
				"genre": genre,
				"page":  page,
			},
			"moviesQuery": moviesQuery,
		})
	})
	// starting and serving
	app.Run("localhost:8080")
	// app.Run(":8080")
}

type PostBody struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

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
