package riki_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"riki/internal/config"
	"time"
)

type RikiClient struct {
	BaseURL string
}

func NewRikiClient(cfg *config.RikiClientConfig) *RikiClient {
	return &RikiClient{BaseURL: cfg.BaseURL}
}
func (c *RikiClient) getByPage(endpoint string, page int, result any) error {
	url := fmt.Sprintf("%s/%s/?page=%d", c.BaseURL, endpoint, page)
	var maxRetries = 5
	var backoff = time.Millisecond * 500

	for i := 0; i <= maxRetries; i++ {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			if i == maxRetries {
				return fmt.Errorf("too many requests, max retries reached")
			}
			log.Printf("retrying in %d milis...", time.Duration(backoff).Milliseconds())
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		var apiError struct {
			Error string `json:"error"`
		}

		if err := json.Unmarshal(body, &apiError); err == nil && apiError.Error != "" {
			if apiError.Error == "There is nothing here" {
				return nil
			}
			return fmt.Errorf("API error: %s", apiError.Error)
		}

		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}

		return nil
	}

	return fmt.Errorf("unexpected error in request loop")
}
