package gateway

import (
	"context"
	"fmt"
	"riki_gateway/internal/model"
	"sync"
)

func (cs *GatewayService) fetchAndSetLocations(ctx context.Context, character *model.Character) error {
	wg := new(sync.WaitGroup)
	errChan := make(chan error, 2)

	var origin, location *model.Location

	if character.Origin != nil {
		wg.Add(1)
		go func(originID int) {
			defer wg.Done()

			res, err := cs.getLocationByID(ctx, originID)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch origin location: %w", err)
				return
			}
			origin = res
		}(*character.Origin.ID)
	}

	if character.Location != nil {
		wg.Add(1)
		go func(locationID int) {
			defer wg.Done()

			res, err := cs.getLocationByID(ctx, locationID)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch character location: %w", err)
				return
			}
			location = res
		}(*character.Location.ID)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	if origin != nil {
		character.Origin = origin
	}
	if location != nil {
		character.Location = location
	}

	return nil
}
