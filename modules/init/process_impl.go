package InitProcess

import (
	"github.com/josevi97/core/logging"
	FileSystemManager "github.com/josevi97/managers/file_system_manager"
)

type InitProcess struct{}

var Log logging.Logging = logging.NewLogging("INIT PROCESS")

func (init *InitProcess) Execute() {
	if !FileSystemManager.Init() {
		Log.Error("could not initialize bluepint")
		Log.Output("blueprint is already initialized")
		return
	}

	Log.Info("executed successfully")
}

func NewInitProcess(args []string) *InitProcess {
	Log.Info("created")
	return &InitProcess{}
}
