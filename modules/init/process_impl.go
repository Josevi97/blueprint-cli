package InitProcess

import (
	"github.com/josevi97/core/logging"

	DatabaseManager "github.com/josevi97/managers/database_manager"
	FileSystemManager "github.com/josevi97/managers/file_system_manager"
	CommandUtils "github.com/josevi97/utils/command"
	ErrorUtils "github.com/josevi97/utils/error"
)

type InitProcess struct {
	pathname string
}

var Log = logging.NewLogging("INIT PROCESS")

func (init *InitProcess) fileSystem() {
	path, err := FileSystemManager.Init(CommandUtils.Pwd())

	if ErrorUtils.IsError(err) {
		ErrorUtils.Throw(err)
	}

	init.pathname = path
}

func (init *InitProcess) database() {
	err := DatabaseManager.Init(init.pathname)

	if ErrorUtils.IsError(err) {
		ErrorUtils.Throw(err)
	}
}

func (init *InitProcess) Execute() {
	init.fileSystem()
	init.database()

	Log.Info("executed successfully")
}

func NewInitProcess(args []string) *InitProcess {
	Log.Info("created")
	return &InitProcess{}
}
