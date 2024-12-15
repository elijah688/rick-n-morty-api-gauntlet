package scan

import (
	"fmt"
	"riki/internal/model"

	"github.com/jackc/pgx/v5"
)

func ScanCharacterEpisode(row pgx.Rows) (*model.CharacterEpisode, error) {
	var cEpisode model.CharacterEpisode

	if err := row.Scan(
		&cEpisode.CharacterID,
		&cEpisode.ID,
		&cEpisode.Name,
		&cEpisode.AirDate,
		&cEpisode.EpisodeCode,
		&cEpisode.URL,
		&cEpisode.Created,
	); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &cEpisode, nil
}
