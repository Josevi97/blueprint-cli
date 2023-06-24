package InitProcess

import (
	logging "github.com/josevi97/core/logging/models"

	CommandUtils "github.com/josevi97/core/command/utils"
	DatabaseManager "github.com/josevi97/core/database/manager"
	FileSystemManager "github.com/josevi97/core/file_system/manager"
)

type InitProcess struct {
	pathname string
}

var Log = logging.NewLogging("INIT PROCESS")

func (initProcess *InitProcess) RequiresToBeInitialized() bool {
	return false
}

func (init *InitProcess) fileSystem() {
	path := FileSystemManager.Init(CommandUtils.Pwd())
	init.pathname = path
}

func (init *InitProcess) database() {
	DatabaseManager.Init(init.pathname)
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
