package crud

import (
	"context"
)

func (cs *CrudService) DeleteCharacter(ctx context.Context, id int) error {

	if err := cs.db.DeleteCharacter(ctx, id); err != nil {
		return err
	}

	return nil
}
