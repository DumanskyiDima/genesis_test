package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/DumanskyiDima/genesis_test/database"
	"github.com/DumanskyiDima/genesis_test/services"
	"github.com/DumanskyiDima/genesis_test/views"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	defer database.Disconnect()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Migrating MongoDB...")
	if err := database.MigrateMongoDB(); err != nil {
		log.Fatalf("Error migrating MongoDB: %v", err)
	}

	// todo: use dictributed task queue like RabbitMQ
	log.Println("Starting cron to send emails...")
	c := cron.New()
	c.AddFunc("* * * * *", services.SendEmails)
	c.Start()

	router := gin.Default()

	router.POST("/subscribe", views.CreateNewSubscription)
	router.GET("/rate", views.GetExchangeRate)

	router.Run("0.0.0.0:8080")
}
