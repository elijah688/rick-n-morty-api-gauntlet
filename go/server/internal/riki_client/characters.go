package riki_client

import "riki/internal/model"

func (c *RikiClient) GetCharactersByPage(page int) ([]model.Character, error) {
	var result struct {
		Results []model.Character `json:"results"`
	}

	if err := c.getByPage("api/character", page, &result); err != nil {
		return nil, err
	}

	return result.Results, nil
}
