package DatabaseManager

import "github.com/josevi97/core/logging"

var Log = logging.NewLogging("DATABASE MANAGER")

func Init(path string) bool {
	Log.Info("initializing")

	return true
}
