package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type ExchangeRate struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	CC           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

// Using cache to reduce amount of requests to the external API
type Cache struct {
	Data      map[string]interface{}
	Timestamp time.Time
}

const cacheTTL = 10 * time.Minute

var (
	cache   Cache
	cacheMu sync.Mutex
)

func FetchExchangeRate() (map[string]interface{}, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	if time.Since(cache.Timestamp) < cacheTTL {
		fmt.Println("Using cached data")
		return cache.Data, nil
	}

	resp, err := http.Get(os.Getenv("EXCHANGE_RATE_API_URL"))
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

	cache.Data = rateDict
	cache.Timestamp = time.Now()

	return rateDict, nil
}
