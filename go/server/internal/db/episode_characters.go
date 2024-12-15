package db

import (
	"context"
	"log"
	"riki/internal/model"
	"sync"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) insertEpisodeCharacters(ctx context.Context, episodes model.Episodes) error {

	records, err := episodes.ToEpisodeCharacterListRecords()
	if err != nil {
		return err
	}
	query, vals, err := goqu.Dialect("postgres").
		Insert("character_episode").
		Rows(records).
		OnConflict(goqu.DoNothing()).
		Prepared(true).
		ToSQL()

	if err != nil {
		return err
	}

	if _, err := db.pool.Exec(ctx, query, vals...); err != nil {
		return err
	}

	log.Printf("successfully inserted %d character_episode\n", len(records))
	return nil
}

func (db *Database) InsertEpisodeCharacters(ctx context.Context, episodes model.Episodes) error {

	wg := new(sync.WaitGroup)
	errChan := make(chan error)

	wg.Add(2)

	go func() {
		defer wg.Done()

	}()

	go func() {
		defer wg.Done()
		if err := db.InsertEpisodes(ctx, episodes); err != nil {
			errChan <- err
			return
		}

		if err := db.insertEpisodeCharacters(ctx, episodes); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}
