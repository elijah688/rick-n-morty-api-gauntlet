package crud

import (
	"context"
)

func (cs *CrudService) GetCharacterEpisodes(ctx context.Context, ids []int) (map[int][]int, error) {

	return cs.db.AllCharacterEpisodesByIDs(ctx, ids)

}
