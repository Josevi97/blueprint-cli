package databases

import (
	database "github.com/josevi97/core/database/models"
	FileSystemUtils "github.com/josevi97/core/file_system/utils"
	logging "github.com/josevi97/core/logging/models"
)

const DATABASE_NAME = "database.db"

var Log = logging.NewLogging("DATABASE")

func DatabaseFactory(path string) database.Database {
	execPath := FileSystemUtils.Join(path, DATABASE_NAME)

	return NewSqlite(execPath)
}
