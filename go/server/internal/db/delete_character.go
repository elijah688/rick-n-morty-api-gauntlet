package db

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) DeleteCharacter(ctx context.Context, characterID int) error {

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	query, vals, err := goqu.
		Dialect("postgres").
		Delete("character_episode").
		Where(goqu.Ex{"character_id": characterID}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build delete query for character_episode: %w", err)
	}

	if _, err := tx.Exec(ctx, query, vals...); err != nil {
		fmt.Println("delete", err)

		return fmt.Errorf("failed to execute delete query for character_episode: %w", err)
	}

	query, vals, err = goqu.
		Dialect("postgres").
		Delete("character").
		Where(goqu.Ex{"id": characterID}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build delete query for character: %w", err)
	}
	if _, err := tx.Exec(ctx, query, vals...); err != nil {
		return fmt.Errorf("failed to execute delete query for character: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
