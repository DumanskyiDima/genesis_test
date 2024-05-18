package main

import (
	"log"

	"github.com/joho/godotenv"

	. "github.com/DumanskyiDima/genesis_test/database"
	. "github.com/DumanskyiDima/genesis_test/views"
	"github.com/gin-gonic/gin"
)

func main() {

	defer Disconnect()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := gin.Default()

	router.POST("/subscribe", CreateNewSubscription)
	router.GET("/rate", GetExchangeRate)

	router.Run("0.0.0.0:8080")
}
