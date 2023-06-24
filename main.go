package main

import (
	"os"

	ErrorHandler "github.com/josevi97/core/error/handler"
	StringUtils "github.com/josevi97/core/string/utils"
	ProcessManager "github.com/josevi97/processes/manager"
)

func main() {
	args := StringUtils.SublistFrom(os.Args, 1)
	process := ProcessManager.FromArgs(args)

	ErrorHandler.Handle(process)
}
