package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) UpsertCharacter(ctx context.Context, char model.Character) (*model.Character, error) {

	c, err := cs.db.UpsertCharacter(ctx, char)

	if err != nil {
		return nil, err
	}

	return c, err
}
