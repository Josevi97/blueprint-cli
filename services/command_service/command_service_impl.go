package services

import "github.com/josevi97/core/database"

type CommandServiceImpl struct {
	repository commandRepository
}

func NewCommandService(db database.Database) CommandService {
	repository := newCommandRepository(db)

	return &CommandServiceImpl{
		repository: repository,
	}
}
