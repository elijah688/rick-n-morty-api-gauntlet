package riki_client

import "riki/internal/model"

func (c *RikiClient) GetEpisodesByPage(page int) ([]model.Episode, error) {
	var result struct {
		Results []model.Episode `json:"results"`
	}
	if err := c.getByPage("api/episode", page, &result); err != nil {
		return nil, err
	}
	return result.Results, nil
}
