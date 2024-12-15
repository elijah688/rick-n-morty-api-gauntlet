package crud

import "riki/internal/db"

type CrudService struct {
	db *db.Database
}

func New(db *db.Database) *CrudService {
	return &CrudService{db}
}
