package process_manager

import (
	"github.com/josevi97/core/logging"
	init_process "github.com/josevi97/src/modules/init"
)

func FromArgs(args []string) Process {
	logging := logging.NewLogging("PROCESS MANAGER")
	logging.Info("Process manager created with %s args", args[0])

	return init_process.NewInitProcess()
}
