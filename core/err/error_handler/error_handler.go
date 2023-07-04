package error_handler

import (
	"github.com/josevi97/core/err"
	"github.com/josevi97/core/err/error_utils"
	"github.com/josevi97/core/file_system/file_system_manager"
	"github.com/josevi97/core/logging"
	"github.com/josevi97/processes"
)

var Log = logging.NewLogging("ERROR HANDLER")

var errorDecider = map[uint]func(){
	err.REQUIRES_INITIALIZATION: func() {
		Log.Error("system must be initialized")
	},
	err.FILESYSTEM_DIRECTRY_ALREADY_EXISTS: func() {
		Log.Error("could not initialize bluepint; already exists")
		Log.Output("blueprint is already initialized")
	},
	err.FILESYSTEM_CREATE_DIRECTORY: func() {
		Log.Error("could not initialize bluepint; could not create directory")
		Log.Output("blueprint is already initialized")
	},
	err.DATABASE_INITIALIZATION: func() {
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

func Handle(proc processes.Process) {
	defer ErrorRecover()

	if proc.RequiresToBeInitialized() && !file_system_manager.IsInitialized() {
		error_utils.Throw(err.REQUIRES_INITIALIZATION)
	}

	proc.Execute()
}
