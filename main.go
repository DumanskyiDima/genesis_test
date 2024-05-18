package main

import (
	"log"

	"github.com/joho/godotenv"

	. "github.com/DumanskyiDima/genesis_test/database"
	. "github.com/DumanskyiDima/genesis_test/services"
	. "github.com/DumanskyiDima/genesis_test/views"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	defer Disconnect()

	// todo: use dictributed task queue like RabbitMQ
	c := cron.New()
	c.AddFunc("51 * * * *", SendEmails)
	c.Start()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Migrating MongoDB...")
	if err := MigrateMongoDB(); err != nil {
		log.Fatalf("Error migrating MongoDB: %v", err)
	}

	router := gin.Default()

	router.POST("/subscribe", CreateNewSubscription)
	router.GET("/rate", GetExchangeRate)

	router.Run("0.0.0.0:8080")
}
