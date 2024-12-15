package db

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (db *Database) AllCharacterEpisodesByIDs(ctx context.Context, characterIDs []int) (map[int][]int, error) {
	sql, vals, err := goqu.
		Dialect("postgres").
		From("character_episode").
		Where(goqu.C("character_id").In(characterIDs)).
		Select(
			"character_id",
			"episode_id",
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

	res := make(map[int][]int, 0)

	for rows.Next() {

		var c, e int

		if err := rows.Scan(&c, &e); err != nil {
			return nil, err
		}

		res[c] = append(res[c], e)

	}

	return res, nil
}
