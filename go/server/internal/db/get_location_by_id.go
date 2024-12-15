package db

import (
	"context"
	"fmt"
	"riki/internal/db/scan"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) GetLocationByID(ctx context.Context, id int) (*model.Location, error) {

	query, vals, err := goqu.Dialect("postgres").
		From("location").
		Select("*").
		Where(goqu.Ex{"id": id}).
		Prepared(true).
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := db.pool.Query(ctx, query, vals...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var locations []model.Location
	for rows.Next() {
		location, err := scan.ScanLocationRow(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)

		}
		locations = append(locations, *location)
	}

	if len(locations) == 0 {
		return nil, fmt.Errorf("no location found with ID %d", id)
	}

	return &locations[0], nil
}
