package services

import database "github.com/josevi97/core/database/models"

const TABLE_NAME = "command"

type commandRepositoryImpl struct {
	repository database.Repository
}

func newCommandRepository(db database.Database) commandRepository {
	repository := database.NewRepository(db, TABLE_NAME)

	return &commandRepositoryImpl{
		repository: repository,
	}
}
