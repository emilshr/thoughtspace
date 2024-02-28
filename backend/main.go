package main

import (
	"backend/common"
	"backend/users"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	common.Init() // Initialize DB connection

	//DB migrations
	users.AutoMigrate()

	router := gin.Default()

	v1 := router.Group("/api/v1")

	//routes
	users.Register(v1.Group("/users"))

	router.Run("localhost:8080")
}
