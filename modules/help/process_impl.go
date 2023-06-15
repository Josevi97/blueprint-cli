package HelpProcess

import (
	"github.com/josevi97/core/logging"
)

type HelpProcess struct{}

var Log logging.Logging = logging.NewLogging("HELP PROCESS")

func (helpProcess *HelpProcess) Execute() {
	Log.Info("Executed")
}

func NewHelpProcess() *HelpProcess {
	Log.Info("Created")
	return &HelpProcess{}
}
