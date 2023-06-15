package main

import (
	"os"

	ProcessManager "github.com/josevi97/core/process_manager"
)

func main() {
	process := ProcessManager.FromArgs(os.Args[1:])
	process.Execute()
}
