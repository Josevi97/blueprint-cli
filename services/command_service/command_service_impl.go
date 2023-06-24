package services

import database "github.com/josevi97/core/database/models"

type CommandServiceImpl struct {
	repository commandRepository
}

func NewCommandService(db database.Database) CommandService {
	repository := newCommandRepository(db)

	return &CommandServiceImpl{
		repository: repository,
	}
}
