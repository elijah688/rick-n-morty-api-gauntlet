package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) UpsertCharacter(ctx context.Context, character model.Character) (*model.Character, error) {
	char, err := cs.upsertCharacter(ctx, character)
	if err != nil {
		return nil, err
	}
	return cs.compileCharacter(ctx, char)

}
func (cs *GatewayService) upsertCharacter(ctx context.Context, character model.Character) (*model.Character, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost
	url := fmt.Sprintf("%s/character", host)

	body, err := json.Marshal(character)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal character: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
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

	var res model.Character
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	return &res, nil
}
