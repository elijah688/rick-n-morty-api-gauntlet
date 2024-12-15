package gateway

import (
	"context"
	"fmt"
	"net/http"
)

func (cs *GatewayService) DeleteCharacter(ctx context.Context, id int) error {
	host := cs.cfg.GatewayConfig.CRUDServiceHost

	url := fmt.Sprintf("%s/character/%d", host, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := cs.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	return nil
}
