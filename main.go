package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"librarymanagement/config"
	"librarymanagement/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
