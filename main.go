package main

import (
	"os"

	ErrorHandler "github.com/josevi97/handlers/error_handler"
	ProcessManager "github.com/josevi97/managers/process_manager"
	StringUtils "github.com/josevi97/utils/string"
)

func main() {
	args := StringUtils.SublistFrom(os.Args, 1)
	process := ProcessManager.FromArgs(args)

	ErrorHandler.Handle(process)
}
