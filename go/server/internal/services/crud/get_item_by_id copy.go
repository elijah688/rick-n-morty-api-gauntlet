package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetLocationByID(ctx context.Context, id int) (*model.Location, error) {

	c, err := cs.db.GetLocationByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return c, err
}
