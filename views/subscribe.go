package views

import (
	"net/http"

	. "github.com/DumanskyiDima/genesis_test/database"

	"github.com/gin-gonic/gin"
)

func CreateNewSubscription(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	err := CreateUser(email, "active")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
