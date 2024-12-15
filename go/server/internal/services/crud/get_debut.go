package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetDebutByID(ctx context.Context, id int) (*model.Episode, error) {

	c, err := cs.db.FirstAppearenceByCharIDs(ctx, []int{id})

	if err != nil {
		return nil, err
	}

	e := c[id]
	return e, nil
}
