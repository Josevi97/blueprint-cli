package ProcessManager

import (
	"strings"

	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/process"
	HelpProcess "github.com/josevi97/modules/help"
	InitProcess "github.com/josevi97/modules/init"
	ArrayUtils "github.com/josevi97/utils/array"
)

var Log = logging.NewLogging("PROCESS MANAGER")

func getProcesses(args []string) map[string]func() process.Process {
	processArgs := ArrayUtils.SublistFrom(args, 1)

	return map[string]func() process.Process{
		process.INIT: func() process.Process {
			return InitProcess.NewInitProcess(processArgs)
		},
		process.HELP: func() process.Process {
			return HelpProcess.NewHelpProcess()
		},
	}
}

func getProcessFromArgs(args []string) process.Process {
	processes := getProcesses(args)
	defaultProcess := processes[process.HELP]

	if args == nil || len(args) == 0 {
		return defaultProcess()
	}

	if process, found := processes[args[0]]; found {
		return process()
	}

	return defaultProcess()
}

func FromArgs(args []string) process.Process {
	Log.Info("Executing process with [%s] arguments", strings.Join(args, ", "))

	return getProcessFromArgs(args)
}
