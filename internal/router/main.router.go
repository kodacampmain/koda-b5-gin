package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
	"github.com/kodacampmain/koda-b5-gin/internal/service"
)

func Init(app *gin.Engine, db *pgxpool.Pool) {
	movieController := controller.NewMovieController()
	rootController := controller.NewRootController()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	// userv2Service := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app.Use(middleware.CORSMiddleware, MyMiddleware)

	app.GET("/", GetRootMiddleware, rootController.GetRoot)
	app.POST("/", rootController.PostRoot)
	app.GET("/movies/:id/:slug", movieController.GetMoviesWithIdAndSlug)
	app.GET("/movies", movieController.SearchAndFilterMoviesWithPagination)
	app.GET("/users", userController.GetUsers)
	app.POST("/users", userController.AddUser)
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
