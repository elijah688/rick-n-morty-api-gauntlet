package gateway

import (
	"context"
	"fmt"
	"riki_gateway/internal/model"
	"sync"
)

func Origins(in []model.Character) []int {
	return extractIDs(in, func(c *model.Character) *int {
		if c.Origin != nil {
			return c.Origin.ID
		}
		return nil
	})
}

func CurrentLocations(in []model.Character) []int {
	return extractIDs(in, func(c *model.Character) *int {
		if c.Location != nil {
			return c.Location.ID
		}
		return nil
	})
}

func extractIDs(in []model.Character, getField func(*model.Character) *int) []int {
	res := make([]int, 0)
	for _, c := range in {
		if id := getField(&c); id != nil {
			res = append(res, *id)
		}
	}
	return res
}

func (cs *GatewayService) fetchAndSetLocations(ctx context.Context, characters []model.Character) ([]model.Character, error) {
	wg := new(sync.WaitGroup)
	errChan := make(chan error, 1)

	var origins, currentLocations map[int]*model.Location

	originsIDs, currentLocationIDs := Origins(characters), CurrentLocations(characters)

	if len(originsIDs) > 0 {
		wg.Add(1)
		go func(oIDs []int) {
			defer wg.Done()

			res, err := cs.getLocationByIDs(ctx, oIDs)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch origin location: %w", err)
				return
			}
			origins = res
		}(originsIDs)
	}

	if len(currentLocationIDs) > 0 {
		wg.Add(1)
		go func(cIDs []int) {
			defer wg.Done()

			res, err := cs.getLocationByIDs(ctx, cIDs)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch character location: %w", err)
				return
			}
			currentLocations = res
		}(currentLocationIDs)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	for i := range characters {
		characters[i].Origin = origins[i]
		characters[i].Location = currentLocations[i]
	}

	return characters, nil
}
