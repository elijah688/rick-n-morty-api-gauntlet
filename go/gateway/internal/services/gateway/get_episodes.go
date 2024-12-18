package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cs *GatewayService) getCharacterEpisodesByIDs(ctx context.Context, ids []int) (map[int][]int, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost
	url := fmt.Sprintf("%s/character/list/episodes", host)

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

	req.Header.Set("Content-Type", "application/json")

	resp, err := cs.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	var episodes map[int][]int
	if err := json.NewDecoder(resp.Body).Decode(&episodes); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return episodes, nil
}
