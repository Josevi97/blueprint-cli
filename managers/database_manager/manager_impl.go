package DatabaseManager

import (
	"github.com/josevi97/core/database"
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"
	"github.com/josevi97/managers/database_manager/databases"
	CommandUtils "github.com/josevi97/utils/command"
)

var Log = logging.NewLogging("DATABASE MANAGER")

type DatabaseManager struct {
	CommandRepository  *database.Repository
	TemplateRepository *database.Repository
}

func Init(path string) uint {
	db := databases.DatabaseFactory(path)

	Log.Info("initializing at %s", path)

	if !db.Open() {
		return error.DATABASE_INITIALIZATION
	}

	db.Migrate()
	db.Close()

	Log.Info("initialized correctly at %s", db.GetPath())

	return error.NONE
}

// TODO: should find a way to make this more abstract

// process => manager => service => repository
// in this way, a only one service would be able to do almost everything
func GetConnection() (*DatabaseManager, uint) {
	path := CommandUtils.Pwd()
	db := databases.DatabaseFactory(path)

	Log.Info("getting connection at %s", path)

	if !db.Open() {
		// TODO: should return a valid error
		return nil, error.NONE
	}

	databaseManager := &DatabaseManager{
		CommandRepository:  &database.Repository{},
		TemplateRepository: &database.Repository{},
	}

	return databaseManager, error.NONE
}
