package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetLocationByID(ctx context.Context, ids []int) (map[int]*model.Location, error) {

	return cs.db.GetLocationByIDs(ctx, ids)

}
