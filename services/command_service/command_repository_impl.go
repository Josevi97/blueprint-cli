package services

import "github.com/josevi97/core/database"

type commandRepositoryImpl struct {
	db *database.Database
}

func newCommandRepository(db *database.Database) commandRepository {
	return &commandRepositoryImpl{
		db: db,
	}
}
