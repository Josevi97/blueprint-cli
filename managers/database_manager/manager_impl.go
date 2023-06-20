package DatabaseManager

import (
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"
	"github.com/josevi97/managers/database_manager/databases"
)

var Log = logging.NewLogging("DATABASE MANAGER")

func Init(path string) uint {
	Log.Info("initializing at %s", path)

	db := databases.DatabaseFactory(path)

	if !db.Open(path) {
		return error.DATABASE_INITIALIZATION
	}

	db.DoMigrations()
	db.Close()

	Log.Info("initialized correctly at %s", db.GetPath())

	return error.NONE
}
