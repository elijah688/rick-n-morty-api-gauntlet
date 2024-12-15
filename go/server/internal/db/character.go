package db

import (
	"context"
	"log"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) InsertCharacters(ctx context.Context, characters model.Characters) error {

	query, vals, err := goqu.Dialect("postgres").
		Insert("character").
		Rows(characters.ToRecords()).
		OnConflict(goqu.DoNothing()).
		Prepared(true).
		ToSQL()

	if err != nil {
		return err
	}

	if _, err := db.pool.Exec(ctx, query, vals...); err != nil {
		return err
	}

	log.Printf("successfully inserted %d characters\n", len(characters))
	return nil
}
