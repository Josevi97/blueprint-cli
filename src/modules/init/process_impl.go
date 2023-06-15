package init_process

import "github.com/josevi97/core/logging"

type InitProcess struct{}

var Log logging.Logging = logging.NewLogging("INIT PROCESS")

func (init *InitProcess) Execute() {
	Log.Info("Executed")
}

func NewInitProcess(args []string) *InitProcess {
	Log.Info("Created")
	return &InitProcess{}
}
