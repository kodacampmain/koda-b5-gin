package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/kodacampmain/koda-b5-gin/docs"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine, db *pgxpool.Pool) {
	app.Use(middleware.CORSMiddleware, MyMiddleware)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterRootRouter(app)
	RegisterMovieRouter(app)
	RegisterUserRouter(app, db)
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
