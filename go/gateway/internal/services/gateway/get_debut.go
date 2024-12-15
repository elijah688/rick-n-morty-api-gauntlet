package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) getDebut(ctx context.Context, id int) (*model.Episode, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost
	url := fmt.Sprintf("%s/character/%d/debut", host, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := cs.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	var episode *model.Episode
	if err := json.NewDecoder(resp.Body).Decode(&episode); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return episode, nil
}
