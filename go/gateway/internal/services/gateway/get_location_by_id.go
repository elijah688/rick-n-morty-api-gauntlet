package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) getLocationByID(ctx context.Context, id int) (*model.Location, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost

	url := fmt.Sprintf("%s/location/%d", host, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	var location model.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &location, nil
}
