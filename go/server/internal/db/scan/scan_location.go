package scan

import (
	"fmt"
	"riki/internal/model"

	"github.com/jackc/pgx/v5"
)

func ScanLocationRow(row pgx.Rows) (*model.Location, error) {
	var location model.Location

	if err := row.Scan(
		&location.ID,
		&location.Name,
		&location.Type,
		&location.Dimension,
		&location.URL,
		&location.Created,
	); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &location, nil
}
