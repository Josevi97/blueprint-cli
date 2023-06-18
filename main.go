package main

import (
	"os"

	ErrorHandler "github.com/josevi97/handlers/error_handler"
	ProcessManager "github.com/josevi97/managers/process_manager"
	ArrayUtils "github.com/josevi97/utils/array"
)

func main() {
	args := ArrayUtils.SublistFrom(os.Args, 1)
	process := ProcessManager.FromArgs(args)

	ErrorHandler.Handle(process.Execute)
}
