package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/controller"
)

func RegisterRootRouter(app *gin.Engine) {
	rootRouter := app.Group("/")

	rootController := controller.NewRootController()

	rootRouter.GET("/", GetRootMiddleware, rootController.GetRoot)
	rootRouter.POST("/", rootController.PostRoot)
}
