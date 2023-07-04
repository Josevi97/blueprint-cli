package command_service

import "github.com/josevi97/core/database"

type CommandServiceImpl struct {
	repository CommandRepository
}

func NewCommandService(db database.Database) CommandService {
	repository := newCommandRepository(db)

	return &CommandServiceImpl{
		repository: repository,
	}
}
