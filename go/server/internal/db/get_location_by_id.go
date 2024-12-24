package db

import (
	"context"
	"fmt"
	"riki/internal/db/scan"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) GetLocationByIDs(ctx context.Context, ids []int) (map[int]*model.Location, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	query, vals, err := goqu.Dialect("postgres").
		From("location").
		Select("*").
		Where(goqu.C("id").In(ids)).
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

	locations := make(map[int]*model.Location, 0)
	for rows.Next() {
		location, err := scan.ScanLocationRow(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)

		}

		locations[*location.ID] = location
	}

	return locations, nil
}
