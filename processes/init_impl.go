package processes

import (
	"github.com/josevi97/core/command/command_utils"
	"github.com/josevi97/core/database/database_manager"
	"github.com/josevi97/core/file_system/file_system_manager"
)

type InitProcess struct {
	pathname string
}

func (initProcess *InitProcess) RequiresToBeInitialized() bool {
	return false
}

func (init *InitProcess) fileSystem() {
	path := file_system_manager.Init(command_utils.Pwd())
	init.pathname = path
}

func (init *InitProcess) database() {
	database_manager.Init(init.pathname)
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
