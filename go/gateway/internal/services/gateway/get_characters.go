package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
	"sync"
)

func (cs *GatewayService) GetCharacters(ctx context.Context, limit, offset int) (*model.CharacterListResponse, error) {

	wg, errChan := new(sync.WaitGroup), make(chan error, 2)
	wg.Add(2)

	chars, cErr := []model.Character(nil), error(nil)
	go func() {
		defer wg.Done()
		chars, cErr = cs.getCharacters(ctx, limit, offset)
		if cErr != nil {
			errChan <- cErr
			return
		}
	}()

	total, tErr := (*int)(nil), error(nil)
	go func() {
		defer wg.Done()

		total, tErr = cs.getTotal(ctx)
		if tErr != nil {
			errChan <- tErr
			return
		}
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(chars)

	return &model.CharacterListResponse{Characters: chars, Total: *total}, nil

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
