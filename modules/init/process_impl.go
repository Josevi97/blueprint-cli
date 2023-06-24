package InitProcess

import (
	"github.com/josevi97/core/logging"

	DatabaseManager "github.com/josevi97/managers/database_manager"
	FileSystemManager "github.com/josevi97/managers/file_system_manager"
	CommandUtils "github.com/josevi97/utils/command"
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
