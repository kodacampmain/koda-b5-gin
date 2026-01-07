package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
)

func RegisterMovieRouter(app *gin.Engine) {
	moviesRouter := app.Group("/movies")

	movieController := controller.NewMovieController()

	moviesRouter.GET("/:id/:slug", movieController.GetMoviesWithIdAndSlug)
	moviesRouter.GET("/", movieController.SearchAndFilterMoviesWithPagination)
}
