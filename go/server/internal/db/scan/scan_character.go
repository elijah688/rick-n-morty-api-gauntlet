package scan

import (
	"fmt"
	"riki/internal/model"

	"github.com/jackc/pgx/v5"
)

func ScanCharacterRow(row pgx.Rows) (*model.Character, error) {
	var character model.Character
	character.Location = new(model.Location)
	character.Origin = new(model.Location)

	if err := row.Scan(
		&character.ID,
		&character.Name,
		&character.Status,
		&character.Species,
		&character.Type,
		&character.Gender,
		&character.Origin.ID,
		&character.Location.ID,
		&character.Image,
		&character.URL,
		&character.Created,
	); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}
	return &character, nil
}
