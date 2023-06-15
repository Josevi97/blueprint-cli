package ProcessManager

import (
	"strings"

	"github.com/josevi97/core/logging"
	HelpProcess "github.com/josevi97/modules/help"
	InitProcess "github.com/josevi97/modules/init"
	ArgsUtils "github.com/josevi97/utils"
)

var Log logging.Logging = logging.NewLogging("PROCESS MANAGER")

func getProcesses(args []string) map[string]func() Process {
	processArgs := ArgsUtils.GetNextArgs(1, args)

	return map[string]func() Process{
		"init": func() Process {
			return InitProcess.NewInitProcess(processArgs)
		},
		"help": func() Process {
			return HelpProcess.NewHelpProcess()
		},
	}
}

func getProcessFromArgs(args []string) Process {
	processes := getProcesses(args)
	defaultProcess := processes["help"]

	if args == nil || len(args) == 0 {
		return defaultProcess()
	}

	if process, found := processes[args[0]]; found {
		return process()
	}

	return defaultProcess()
}

func FromArgs(args []string) Process {
	Log.Info("Executing process with [%s] arguments", strings.Join(args, ", "))

	return getProcessFromArgs(args)
}
