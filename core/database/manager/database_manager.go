package Database

import (
	databases "github.com/josevi97/core/database/factory"
	database "github.com/josevi97/core/database/models"
	error "github.com/josevi97/core/error/models"
	ErrorUtils "github.com/josevi97/core/error/utils"
	logging "github.com/josevi97/core/logging/models"
)

var Log = logging.NewLogging("DATABASE MANAGER")

func Init(path string) {
	db := GetConnection(path)

	Log.Info("initializing at %s", path)

	if !db.Open() {
		ErrorUtils.Throw(error.DATABASE_INITIALIZATION)
	}

	db.Migrate()
	db.Close()

	Log.Info("initialized correctly at %s", db.GetPath())
}

// TODO: Needs to be tested. Looks like it will need to be a pointer.
// However, may not be necessary as soon as Database is only an interface
func GetConnection(path string) database.Database {
	db := databases.DatabaseFactory(path)

	Log.Info("getting connection at %s", path)

	return db
}
