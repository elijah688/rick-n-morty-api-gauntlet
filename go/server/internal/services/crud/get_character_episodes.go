package crud

import (
	"context"
)

func (cs *CrudService) GetCharacterEpisodes(ctx context.Context, id int) ([]int, error) {

	mEs, err := cs.db.AllCharacterEpisodesByIDs(ctx, []int{id})

	if err != nil {
		return nil, err
	}

	es := mEs[id]

	return es, nil
}
