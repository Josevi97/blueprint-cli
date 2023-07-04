package file_system_manager

import (
	"github.com/josevi97/core/command/command_utils"
	"github.com/josevi97/core/err"
	"github.com/josevi97/core/err/error_utils"
	"github.com/josevi97/core/file_system/file_system_utils"
	"github.com/josevi97/core/logging"
)

const DIRECTORY_NAME = ".blueprint"

var Log = logging.NewLogging("FILE SYSTEM MANAGER")

func Init(path string) string {
	execPath := file_system_utils.Join(path, DIRECTORY_NAME)

	Log.Info("initializing at %s", path)

	if file_system_utils.Exists(DIRECTORY_NAME, path) {
		error_utils.Throw(err.FILESYSTEM_DIRECTRY_ALREADY_EXISTS)
	}

	if !file_system_utils.CreateDirectory(DIRECTORY_NAME, path) {
		error_utils.Throw(err.FILESYSTEM_CREATE_DIRECTORY)
	}

	Log.Info("initialized correctly at \"%s\"", execPath)

	return execPath
}

func GetCurrentPath() string {
	return file_system_utils.Join(command_utils.Pwd(), DIRECTORY_NAME)
}

func IsInitialized() bool {
	return file_system_utils.Exists(DIRECTORY_NAME, command_utils.Pwd())
}
