package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/router"
)

func main() {
	// initialization
	app := gin.Default()
	// routing
	router.Init(app)
	// starting and serving
	app.Run("localhost:8080")
	// app.Run(":8080")
}
