package databases

import (
	"github.com/josevi97/core/database"
	"github.com/josevi97/core/logging"
	FileSystemUtils "github.com/josevi97/utils/file_system"
)

const DATABASE_NAME = "database.db"

var Log = logging.NewLogging("DATABASE")

func DatabaseFactory(path string) database.Database {
	execPath := FileSystemUtils.Join(path, DATABASE_NAME)

	return NewSqlite(execPath)
}
