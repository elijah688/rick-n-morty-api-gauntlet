package riki_client

import "riki/internal/model"

func (c *RikiClient) GetLocationsByPage(page int) ([]model.Location, error) {
	var result struct {
		Results []model.Location `json:"results"`
	}

	if err := c.getByPage("api/location", page, &result); err != nil {
		return nil, err
	}

	return result.Results, nil
}
