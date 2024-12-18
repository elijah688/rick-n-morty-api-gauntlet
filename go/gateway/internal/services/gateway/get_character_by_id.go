package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) getCharacterByID(ctx context.Context, id int) (*model.Character, error) {
	host := cs.cfg.GatewayConfig.CRUDServiceHost

	url := fmt.Sprintf("%s/character/%d", host, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var character model.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &character, nil
}
