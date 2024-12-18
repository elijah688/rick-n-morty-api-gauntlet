package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) getLocationByIDs(ctx context.Context, ids []int) (map[int]*model.Location, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost

	url := fmt.Sprintf("%s/location/list", host)

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

	var location map[int]*model.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return location, nil
}

// # curl "http://localhost:8080/character/list/episodes" \
// #     -H "content-type: application/json" \
// #     -d '{"ids":[1,2,3]}' | jq .

// # curl "http://localhost:8080/location/list" \
// # -H "content-type: application/json" \
// # -d '{"ids":[1,2,3]}' | jq .

// curl "http://localhost:8080/character/list/debut" \
//     -H "content-type: application/json" \
//     -d '{"ids":[1,2,3]}' | jq .
