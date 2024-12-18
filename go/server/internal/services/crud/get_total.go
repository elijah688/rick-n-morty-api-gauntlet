package crud

import (
	"context"
)

func (cs *CrudService) GetTotal(ctx context.Context) (int, error) {

	return cs.db.GetTotal(ctx)

}
