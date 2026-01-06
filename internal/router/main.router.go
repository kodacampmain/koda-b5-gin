package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
)

func Init(app *gin.Engine) {
	movieController := controller.NewMovieController()
	rootController := controller.NewRootController()

	app.GET("/", rootController.GetRoot)
	app.POST("/", rootController.PostRoot)
	app.GET("/movies/:id/:slug", movieController.GetMoviesWithIdAndSlug)
	app.GET("/movies", movieController.SearchAndFilterMoviesWithPagination)
}
