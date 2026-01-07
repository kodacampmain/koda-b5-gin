package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kodacampmain/koda-b5-gin/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to Load env")
		return
	}
	// log.Println(os.Getenv("DB_HOST"))
	// setup db
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, pwd, host, port, dbName)
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Println("Failed to Connect to Database")
		return
	}
	// initialization
	app := gin.Default()
	// routing
	router.Init(app, db)
	// starting and serving
	app.Run("localhost:8080")
	// app.Run(":8080")
}
