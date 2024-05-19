package views

import (
	"log"
	"net/http"

	"github.com/DumanskyiDima/genesis_test/database"

	"github.com/gin-gonic/gin"
)

func CreateNewSubscription(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	existing_user := database.FindUser(email)
	if existing_user != nil {
		if existing_user.Status == "active" {
			c.JSON(http.StatusConflict, gin.H{"message": "Already subscribed"})
			return
		} else {
			// todo make user active
			log.Println("User already exists but not active")
		}
	}

	_, err := database.CreateUser(email, "active")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User subscribed successfully"})
}
