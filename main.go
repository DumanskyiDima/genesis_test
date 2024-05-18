package main

import (
	. "github.com/DumanskyiDima/genesis_test/database"
	. "github.com/DumanskyiDima/genesis_test/views"
	"github.com/gin-gonic/gin"
)

func main() {
	defer Disconnect()
	router := gin.Default()

	router.POST("/subscribe", CreateNewSubscription)
	router.GET("/rate", GetExchangeRate)

	router.Run("localhost:8080")
}
