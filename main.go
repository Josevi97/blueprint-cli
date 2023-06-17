package main

import (
	"os"

	ProcessManager "github.com/josevi97/managers/process_manager"
	ArrayUtils "github.com/josevi97/utils/array"
)

func main() {
	args := ArrayUtils.SublistFrom(os.Args, 1)
	process := ProcessManager.FromArgs(args)

	process.Execute()
}
