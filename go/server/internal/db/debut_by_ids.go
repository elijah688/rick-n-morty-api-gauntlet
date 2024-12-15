package db

import (
	"context"
	"fmt"
	"riki/internal/db/scan"
	"riki/internal/model"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) FirstAppearenceByCharIDs(ctx context.Context, characterIDs []int) (map[int]*model.Episode, error) {
	sql, vals, err := goqu.Dialect("postgres").
		From("ranked_episodes").
		Where(goqu.C("rn").Eq(1)).
		Order(goqu.I("character_id").Asc()).
		With(
			"ranked_episodes",
			goqu.Dialect("postgres").
				Select(
					goqu.I("ce.character_id"),
					goqu.I("ce.episode_id"),
					goqu.I("e.name").As("episode_name"),
					goqu.I("e.air_date"),
					goqu.I("e.episode_code"),
					goqu.I("e.url").As("episode_url"),
					goqu.I("e.created").As("episode_created"),
					goqu.L("ROW_NUMBER() OVER (PARTITION BY ce.character_id ORDER BY ce.episode_id ASC)").As("rn"),
				).
				From(goqu.I("character_episode").As("ce")).
				Join(
					goqu.I("episode").As("e"),
					goqu.On(goqu.Ex{"ce.episode_id": goqu.I("e.id")}),
				).
				Where(
					goqu.I("ce.character_id").In(characterIDs),
				),
		).
		Select(
			goqu.I("character_id"),
			goqu.I("episode_id").As("earliest_episode_id"),
			goqu.I("episode_name"),
			goqu.I("air_date"),
			goqu.I("episode_code"),
			goqu.I("episode_url"),
			goqu.I("episode_created"),
		).
		Prepared(true).
		ToSQL()

	if err != nil {
		return nil, err
	}

	rows, err := db.pool.Query(ctx, sql, vals...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	characterEpisodes := make(map[int]*model.Episode)

	for rows.Next() {

		ce, err := scan.ScanCharacterEpisode(rows)
		if err != nil {
			return nil, err
		}

		characterEpisodes[ce.CharacterID] = &ce.Episode

	}

	return characterEpisodes, nil
}
