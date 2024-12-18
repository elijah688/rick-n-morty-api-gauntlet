package gateway

import (
	"context"
	"riki_gateway/internal/model"
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

	locations, err := cs.getLocationByIDs(ctx, append(Origins(characters), CurrentLocations(characters)...))
	if err != nil {
		return nil, err
	}

	for i := range characters {
		c := &characters[i]

		if c.Origin != nil && c.Origin.ID != nil {
			if loc, ok := locations[*c.Origin.ID]; ok {
				c.Origin = loc
			}
		}

		if c.Location != nil && c.Location.ID != nil {
			if loc, ok := locations[*c.Location.ID]; ok {
				c.Location = loc
			}
		}
	}

	return characters, nil
}
