package views

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ExchangeRate struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	CC           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

// Using cache to reduce the number of requests to the API
type Cache struct {
	Data      map[string]interface{}
	Timestamp time.Time
}

const (
	url      = "https://bank.gov.ua/NBUStatService/v1/statdirectory/dollar_info?json"
	cacheTTL = 10 * time.Minute
)

var (
	cache   Cache
	cacheMu sync.Mutex
)

func fetchExchangeRate() (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var rates []ExchangeRate
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	if len(rates) == 0 {
		return nil, fmt.Errorf("no exchange rate data found")
	}

	rateDict := map[string]interface{}{
		"r030":         rates[0].R030,
		"txt":          rates[0].Txt,
		"rate":         rates[0].Rate,
		"cc":           rates[0].CC,
		"exchangedate": rates[0].ExchangeDate,
	}

	return rateDict, nil
}

func GetExchangeRate(c *gin.Context) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	var rate map[string]interface{}
	var err error

	if time.Since(cache.Timestamp) < cacheTTL {
		rate, err = cache.Data, nil
		fmt.Println("Using cached data")

	} else {
		rate, err = fetchExchangeRate()
		cache.Data = rate
		cache.Timestamp = time.Now()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rate)
}
