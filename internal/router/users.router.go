package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
	"github.com/kodacampmain/koda-b5-gin/internal/service"
	"github.com/redis/go-redis/v9"
)

func RegisterUserRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	userRouter := app.Group("/users")

	userRepository := repository.NewUserRepository(db)
	// userRepository := repository.NewUserRepositoryMock()
	userService := service.NewUserService(userRepository, rdb)
	// userv2Service := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	userRouter.GET("", userController.GetUsers)
	// userRouter.POST("/", userController.AddUser)
	userRouter.PATCH("", middleware.VerifyJWT, middleware.UserOnly, userController.EditProfile)

	userRouter.POST("/auth/new", userController.AddUser)
	userRouter.POST("/auth", userController.Login)

}
