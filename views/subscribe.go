package views

import (
	"net/http"

	. "github.com/DumanskyiDima/genesis_test/database"

	"github.com/gin-gonic/gin"
)

func CreateNewSubscription(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	user := FindUser(email)
	if user != nil {
		if user.Status == "active" {
			c.JSON(http.StatusConflict, gin.H{"message": "Already subscribed"})
			return
		}
	}

	err := CreateUser(email, "active")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User subscribed successfully"})
}
