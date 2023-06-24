package FileSystemUtils

import (
	"io/fs"
	"os"

	StringUtils "github.com/josevi97/core/string/utils"
)

func Join(data ...string) string {
	return StringUtils.Join(data, "/")
}

func Exists(name string, path string) bool {
	_, err := os.Stat(Join(path, name))

	if err == nil {
		return true
	}

	if os.IsExist(err) {
		return true
	}

	// This may be an error anyways. This should be refactored to use error status
	// or panic it
	return false
}

func CreateDirectory(name string, path string) bool {
	directoryName := Join(path, name)
	err := os.Mkdir(directoryName, fs.ModePerm)

	if err != nil {
		return false
	}

	return true
}
