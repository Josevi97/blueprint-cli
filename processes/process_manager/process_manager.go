package process_manager

import (
	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/string/string_utils"
	"github.com/josevi97/processes"
)

var Log = logging.NewLogging("PROCESS MANAGER")

func getProcesses(args []string) map[string]func() processes.Process {
	processArgs := string_utils.SublistFrom(args, 1)

	return map[string]func() processes.Process{
		processes.INIT: func() processes.Process {
			return processes.NewInitProcess(processArgs)
		},
		processes.HELP: func() processes.Process {
			return processes.NewHelpProcess()
		},
		processes.REGISTER: func() processes.Process {
			return processes.NewRegisterProcess()
		},
	}
}

func getProcessFromArgs(args []string) processes.Process {
	pro := getProcesses(args)
	defaultProcess := pro[processes.HELP]

	if args == nil || len(args) == 0 {
		return defaultProcess()
	}

	if process, found := pro[args[0]]; found {
		return process()
	}

	return defaultProcess()
}

func FromArgs(args []string) processes.Process {
	Log.Info("Executing process with [%s] arguments", string_utils.Join(args, ", "))

	return getProcessFromArgs(args)
}
