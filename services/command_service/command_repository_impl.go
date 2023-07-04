package command_service

import "github.com/josevi97/core/database"

const TABLE_NAME = "command"

type commandRepositoryImpl struct {
	repository database.Repository
}

func newCommandRepository(db database.Database) CommandRepository {
	repository := database.NewRepository(db, TABLE_NAME)

	return &commandRepositoryImpl{
		repository: repository,
	}
}
