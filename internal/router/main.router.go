package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/kodacampmain/koda-b5-gin/docs"
	"github.com/kodacampmain/koda-b5-gin/internal/middleware"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.Use(middleware.CORSMiddleware, MyMiddleware)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Static("/static/img", "public")
	app.Static("/static/pages", "public/html")

	RegisterRootRouter(app)
	RegisterMovieRouter(app)
	RegisterUserRouter(app, db, rdb)

	app.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/static/pages/not-found.html")
	})
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
