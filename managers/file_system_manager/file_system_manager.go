package FileSystemManager

import (
	"github.com/josevi97/core/logging"

	CommandUtils "github.com/josevi97/utils/command"
	FileSystemUtils "github.com/josevi97/utils/file_system"
)

var Log logging.Logging = logging.NewLogging("FILE SYSTEM MANAGER")

func Init() bool {
	execPath := CommandUtils.Pwd()
	directoryName := FileSystemUtils.DIRECTORY_NAME

	Log.Info("initializing at %s", execPath)

	if FileSystemUtils.Exists(directoryName, execPath) {
		Log.Error("directory already exists")
		return false
	}

	if !FileSystemUtils.CreateDirectory(directoryName, execPath) {
		Log.Error("could not create blueprint directory")
		return false
	}

	Log.Info("initialized correctly \"%s\"", FileSystemUtils.Join(execPath, directoryName))

	return true
}
