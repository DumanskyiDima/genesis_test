package views

import (
	"net/http"

	"github.com/DumanskyiDima/genesis_test/services"

	"github.com/gin-gonic/gin"
)

func GetExchangeRate(c *gin.Context) {
	rate, err := services.FetchExchangeRate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rate)
}
