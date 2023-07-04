package database_factory

import (
	"github.com/josevi97/core/database"
	"github.com/josevi97/core/file_system/file_system_utils"
	"github.com/josevi97/core/logging"
)

const DATABASE_NAME = "database.db"

var Log = logging.NewLogging("DATABASE")

func NewDatabase(path string) database.Database {
	execPath := file_system_utils.Join(path, DATABASE_NAME)

	return NewSqlite(execPath)
}
