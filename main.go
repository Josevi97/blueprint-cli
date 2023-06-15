package main

import (
	"os"

	"github.com/josevi97/core/process_manager"
)

func main() {
	process_manager.FromArgs(os.Args[1:]).Execute()
}
