package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kodacampmain/koda-b5-gin/internal/config"
	"github.com/kodacampmain/koda-b5-gin/internal/router"
)

// @title           Koda 5 Gin
// @version         1.0
// @description     Backend class using go with gin
// @host      		localhost:8080
// @BasePath  		/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to Load env")
		return
	}
	// db initialization
	db, err := config.InitDb()
	if err != nil {
		log.Println("Failed to Connect to Database")
		return
	}
	// gin initialization
	app := gin.Default()
	// routing
	router.Init(app, db)
	// starting and serving
	app.Run("localhost:8080")
	// app.Run(":8080")
}
