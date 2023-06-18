package DatabaseManager

import (
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"
)

var Log = logging.NewLogging("DATABASE MANAGER")

func Init(path string) uint {
	Log.Info("initializing")

	return error.NONE
}
