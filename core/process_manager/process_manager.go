package process_manager

import (
	"strings"

	"github.com/josevi97/core/logging"
	help_process "github.com/josevi97/src/modules/help"
	init_process "github.com/josevi97/src/modules/init"
)

var Log logging.Logging = logging.NewLogging("PROCESS MANAGER")

func getProcessArgs(args []string) []string {
	if len(args) == 1 {
		return []string{}
	}

	return args[1:]
}

func getProcesses(args []string) map[string]func() Process {
	processArgs := getProcessArgs(args)

	return map[string]func() Process{
		"init": func() Process {
			return init_process.NewInitProcess(processArgs)
		},
		"help": func() Process {
			return help_process.NewHelpProcess()
		},
	}
}

func getProcessFromArgs(args []string) Process {
	if args == nil || len(args) == 0 {
		return help_process.NewHelpProcess()
	}

	processes := getProcesses(args)
	process, found := processes[args[0]]

	if !found {
		return processes["help"]()
	}

	return process()
}

func FromArgs(args []string) Process {
	Log.Info("Executing process with [%s] arguments", strings.Join(args, ", "))

	return getProcessFromArgs(args)
}
