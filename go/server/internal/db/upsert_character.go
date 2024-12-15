package db

import (
	"context"
	"fmt"
	"riki/internal/db/scan"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
)

func doUpdate() exp.ConflictUpdateExpression {
	return goqu.DoUpdate(
		"id",
		goqu.Record{
			"name":        goqu.L("EXCLUDED.name"),
			"status":      goqu.L("EXCLUDED.status"),
			"species":     goqu.L("EXCLUDED.species"),
			"type":        goqu.L("EXCLUDED.type"),
			"gender":      goqu.L("EXCLUDED.gender"),
			"origin_id":   goqu.L("EXCLUDED.origin_id"),
			"location_id": goqu.L("EXCLUDED.location_id"),
			"image":       goqu.L("EXCLUDED.image"),
			"url":         goqu.L("EXCLUDED.url"),
			"created":     goqu.L("EXCLUDED.created"),
		},
	)
}

func (db *Database) UpsertCharacter(ctx context.Context, character model.Character) (*model.Character, error) {

	query, vals, err := goqu.
		Dialect("postgres").
		Insert("character").
		Rows(character.ToRecord()).
		OnConflict(doUpdate()).
		Returning("*").
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build upsert query: %w", err)
	}

	rows, err := db.pool.Query(ctx, query, vals...)
	if err != nil {
		return nil, fmt.Errorf("failed upserting char: %w", err)
	}

	res, err := (*model.Character)(nil), error(nil)
	for rows.Next() {
		res, err = scan.ScanCharacterRow(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan upserted character: %w", err)
		}
	}

	return res, nil
}
