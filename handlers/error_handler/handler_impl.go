package ErrorHandler

import (
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/process"
	FileSystemManager "github.com/josevi97/managers/file_system_manager"
	ErrorUtils "github.com/josevi97/utils/error"
)

var Log = logging.NewLogging("ERROR HANDLER")

var errorDecider = map[uint]func(){
	error.REQUIRES_INITIALIZATION: func() {
		Log.Error("system must be initialized")
	},
	error.FILESYSTEM_DIRECTRY_ALREADY_EXISTS: func() {
		Log.Error("could not initialize bluepint; already exists")
		Log.Output("blueprint is already initialized")
	},
	error.FILESYSTEM_CREATE_DIRECTORY: func() {
		Log.Error("could not initialize bluepint; could not create directory")
		Log.Output("blueprint is already initialized")
	},
	error.DATABASE_INITIALIZATION: func() {
		Log.Error("could not initialize the database")
	},
}

func ErrorRecover() {
	var err interface{} = recover()

	if err == nil {
		return
	}

	if key, ok := err.(uint); ok {
		errorDecider[key]()
	} else {
		Log.Error("unknown system error throwed: %s", err)
	}
}

func Handle(proc process.Process) {
	defer ErrorRecover()

	if proc.RequiresToBeInitialized() && !FileSystemManager.IsInitialized() {
		ErrorUtils.Throw(error.REQUIRES_INITIALIZATION)
	}

	proc.Execute()
}
