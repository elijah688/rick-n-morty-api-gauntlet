package db

import (
	"context"
	"fmt"
	"riki/internal/db/scan"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) GetCharacters(ctx context.Context, limit, offset int) ([]model.Character, error) {

	query, vals, err := goqu.
		Dialect("postgres").
		From("character").
		Select("*").
		Limit(uint(limit)).
		Offset(uint(offset)).
		Order(goqu.C("id").Asc()).
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

	var characters []model.Character
	for rows.Next() {
		character, err := scan.ScanCharacterRow(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)

		}
		characters = append(characters, *character)
	}

	return characters, nil
}

func (db *Database) GetTotal(ctx context.Context) (int, error) {
	query, _, err := goqu.
		Dialect("postgres").
		From("character").
		Select(goqu.L("COUNT(*) AS total")).
		Prepared(true).
		ToSQL()

	if err != nil {
		return 0, fmt.Errorf("failed to build count query: %w", err)
	}

	var totalCount int
	err = db.pool.QueryRow(ctx, query).Scan(&totalCount)
	if err != nil {
		return 0, fmt.Errorf("failed to execute count query: %w", err)
	}

	return totalCount, nil
}
