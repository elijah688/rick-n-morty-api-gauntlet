package db

import (
	"context"
	"log"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) InsertEpisodes(ctx context.Context, episode model.Episodes) error {

	query, vals, err := goqu.Dialect("postgres").
		Insert("episode").
		Rows(episode.ToRecords()).
		OnConflict(goqu.DoNothing()).
		Prepared(true).
		ToSQL()

	if err != nil {
		return err
	}

	if _, err := db.pool.Exec(ctx, query, vals...); err != nil {
		return err
	}

	log.Printf("successfully inserted %d episodes\n", len(episode))
	return nil
}
