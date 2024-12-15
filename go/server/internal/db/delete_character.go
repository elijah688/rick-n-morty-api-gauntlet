package db

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) DeleteCharacter(ctx context.Context, characterID int) error {
	query, vals, err := goqu.
		Dialect("postgres").
		Delete("character").
		Where(goqu.Ex{"id": characterID}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	if _, err := db.pool.Exec(ctx, query, vals...); err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	return nil
}
