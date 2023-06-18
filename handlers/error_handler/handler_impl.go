package ErrorHandler

import (
	"github.com/josevi97/core/error"
	"github.com/josevi97/core/logging"
)

var Log = logging.NewLogging("ERROR HANDLER")

var errorDecider = map[uint]func(){
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
		Log.Error("unknown error throwed")
	}
}

func Handle(process func()) {
	defer ErrorRecover()

	process()
}
