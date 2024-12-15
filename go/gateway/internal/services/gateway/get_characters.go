package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) GetCharacters(ctx context.Context, limit, offset int) ([]model.Character, error) {
	res, err := cs.getCharacters(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range res {
		c := res[i]
		r, err := cs.compileCharacter(ctx, &c)
		if err != nil {
			return nil, err
		}
		res[i] = *r
	}

	return res, nil
}

func (cs *GatewayService) getCharacters(ctx context.Context, limit, offset int) ([]model.Character, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost

	url := fmt.Sprintf("%s/character?limit=%d&offset=%d", host, limit, offset)

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

	var res []model.Character
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return res, nil
}
