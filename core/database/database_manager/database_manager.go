package database_manager

import (
	"github.com/josevi97/core/database"
	"github.com/josevi97/core/database/database_factory"
	"github.com/josevi97/core/err"
	"github.com/josevi97/core/err/error_utils"
	"github.com/josevi97/core/logging"
)

var Log = logging.NewLogging("DATABASE MANAGER")

func Init(path string) {
	db := GetConnection(path)

	Log.Info("initializing at %s", path)

	if !db.Open() {
		error_utils.Throw(err.DATABASE_INITIALIZATION)
	}

	db.Migrate()
	db.Close()

	Log.Info("initialized correctly at %s", db.GetPath())
}

// TODO: Needs to be tested. Looks like it will need to be a pointer.
// However, may not be necessary as soon as Database is only an interface
func GetConnection(path string) database.Database {
	db := database_factory.NewDatabase(path)

	Log.Info("getting connection at %s", path)

	return db
}
