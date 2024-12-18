package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) getDebusByIDs(ctx context.Context, ids []int) (map[int]*model.Episode, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost
	url := fmt.Sprintf("%s/character/list/debut", host)

	bodyReader, err := marshalToJSONBody(
		map[string]any{
			"ids": ids,
		})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
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

	var episode map[int]*model.Episode
	if err := json.NewDecoder(resp.Body).Decode(&episode); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return episode, nil
}
