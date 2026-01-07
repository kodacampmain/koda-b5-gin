package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
)

func Init(app *gin.Engine) {
	movieController := controller.NewMovieController()
	rootController := controller.NewRootController()

	app.Use(middleware.CORSMiddleware, MyMiddleware)

	app.GET("/", GetRootMiddleware, rootController.GetRoot)
	app.POST("/", rootController.PostRoot)
	app.GET("/movies/:id/:slug", movieController.GetMoviesWithIdAndSlug)
	app.GET("/movies", movieController.SearchAndFilterMoviesWithPagination)
}

func MyMiddleware(c *gin.Context) {
	// sebelum controller di alur request
	// request -> middleware -> controller
	log.Println("BEFORE")
	c.Next()
	log.Println("AFTER")
	// sesudah controller di alur response
	// controller -> middleware -> response
}

func GetRootMiddleware(c *gin.Context) {
	defer c.Next()
	log.Printf("HOST: %s\n", c.GetHeader("Origin"))
}
