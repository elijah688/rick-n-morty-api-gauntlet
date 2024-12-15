package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetCharacterByID(ctx context.Context, id int) (*model.Character, error) {

	c, err := cs.db.GetCharacterByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return c, err
}
