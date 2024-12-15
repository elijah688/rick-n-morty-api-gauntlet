package services

import (
	"riki/internal/db"
	"riki/internal/services/crud"
)

type Services struct {
	crud *crud.CrudService
}

func New(db *db.Database) *Services {
	return &Services{crud: crud.New(db)}
}

func (svcs *Services) CRUD() *crud.CrudService {
	return svcs.crud
}
