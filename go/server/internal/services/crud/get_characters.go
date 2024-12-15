package crud

import (
	"context"
	"riki/internal/model"
)

func (cs *CrudService) GetCharacters(ctx context.Context, limit, offset int) ([]model.Character, error) {

	chars, err := cs.db.GetCharacters(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	return chars, err
}
