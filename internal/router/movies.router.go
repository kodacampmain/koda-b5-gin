package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
)

func RegisterMovieRouter(app *gin.Engine) {
	moviesRouter := app.Group("/movies")

	movieController := controller.NewMovieController()

	moviesRouter.Use(middleware.VerifyJWT)

	moviesRouter.GET("/:id/:slug", middleware.AdminOnly, movieController.GetMoviesWithIdAndSlug)
	moviesRouter.GET("/", middleware.UserOnly, movieController.SearchAndFilterMoviesWithPagination)
}
