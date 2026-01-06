package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
)

type MovieController struct{}

func NewMovieController() *MovieController {
	return &MovieController{}
}

func (m *MovieController) GetMoviesWithIdAndSlug(c *gin.Context) {
	// direct access
	id, _ := strconv.Atoi(c.Param("id"))
	slug := c.Param("slug")
	// data binding
	var moviesParam dto.MoviesParam
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
}

func (m *MovieController) SearchAndFilterMoviesWithPagination(c *gin.Context) {
	// direct access
	title := c.Query("title")
	genre := c.QueryArray("genre")
	page := c.Query("page")
	// data binding
	var moviesQuery dto.MoviesQuery
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
}
