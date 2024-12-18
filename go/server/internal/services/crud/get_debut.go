package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetDebutByIDs(ctx context.Context, ids []int) (map[int]*model.Episode, error) {

	return cs.db.FirstAppearenceByCharIDs(ctx, ids)

}
