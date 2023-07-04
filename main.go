package main

import (
	"os"

	"github.com/josevi97/core/err/error_handler"
	"github.com/josevi97/core/string/string_utils"
	"github.com/josevi97/processes/process_manager"
)

func main() {
	args := string_utils.SublistFrom(os.Args, 1)
	process := process_manager.FromArgs(args)

	error_handler.Handle(process)
}
