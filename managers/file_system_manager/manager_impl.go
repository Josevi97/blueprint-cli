package FileSystemManager

import (
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"

	FileSystemUtils "github.com/josevi97/utils/file_system"
)

const DIRECTORY_NAME = ".blueprint"

var Log = logging.NewLogging("FILE SYSTEM MANAGER")

func Init(path string) (string, uint) {
	execPath := FileSystemUtils.Join(path, DIRECTORY_NAME)

	Log.Info("initializing at %s", path)

	if FileSystemUtils.Exists(DIRECTORY_NAME, path) {
		return "", error.FILESYSTEM_DIRECTRY_ALREADY_EXISTS
	}

	if !FileSystemUtils.CreateDirectory(DIRECTORY_NAME, path) {
		return "", error.FILESYSTEM_CREATE_DIRECTORY
	}

	Log.Info("initialized correctly at \"%s\"", execPath)

	return execPath, error.NONE
}
