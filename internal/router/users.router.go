package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
	"github.com/kodacampmain/koda-b5-gin/internal/service"
)

func RegisterUserRouter(app *gin.Engine, db *pgxpool.Pool) {
	userRouter := app.Group("/users")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	// userv2Service := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	userRouter.GET("/", userController.GetUsers)
	userRouter.POST("/", userController.AddUser)

	userRouter.POST("/auth/new", userController.Register)
	userRouter.POST("/auth", userController.Login)
}
