package FileSystem

import (
	error "github.com/josevi97/core/error/models"
	logging "github.com/josevi97/core/logging/models"

	CommandUtils "github.com/josevi97/core/command/utils"
	ErrorUtils "github.com/josevi97/core/error/utils"
	FileSystemUtils "github.com/josevi97/core/file_system/utils"
)

const DIRECTORY_NAME = ".blueprint"

var Log = logging.NewLogging("FILE SYSTEM MANAGER")

func Init(path string) string {
	execPath := FileSystemUtils.Join(path, DIRECTORY_NAME)

	Log.Info("initializing at %s", path)

	if FileSystemUtils.Exists(DIRECTORY_NAME, path) {
		ErrorUtils.Throw(error.FILESYSTEM_DIRECTRY_ALREADY_EXISTS)
	}

	if !FileSystemUtils.CreateDirectory(DIRECTORY_NAME, path) {
		ErrorUtils.Throw(error.FILESYSTEM_CREATE_DIRECTORY)
	}

	Log.Info("initialized correctly at \"%s\"", execPath)

	return execPath
}

func GetCurrentPath() string {
	return FileSystemUtils.Join(CommandUtils.Pwd(), DIRECTORY_NAME)
}

func IsInitialized() bool {
	return FileSystemUtils.Exists(DIRECTORY_NAME, CommandUtils.Pwd())
}
