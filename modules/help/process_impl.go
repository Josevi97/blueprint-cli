package HelpProcess

import (
	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/process"
)

type HelpProcess struct{}

var Log = logging.NewLogging("HELP PROCESS")

func args(command string, msg string) {
	Log.Output("%s\t\t%s", command, msg)
}

func (helpProcess *HelpProcess) Execute() {
	Log.Output("Blueprint CLI for structuring your project defining conventions.")
	Log.Output("The basic CLI commands are as follows:\n")

	args(process.HELP, "\tdisplays the blueprint CLI commands available\n")
	args(process.INIT, "\tuse this command to initialize blueprint CLI\n")
	args(process.REGISTER, "registers a block structure in order to get a command to generate a block item")
	args("  -c --command", "dictamines the command to be used for generate command\n")

	args(process.GENERATE, "generates a block item depending on the blocks registered\n")
	args(process.LIST, "\tlist every command block generated\n")
	args(process.BLUEPRINT, "opens GUI blueprint\n")

	Log.Output("\nCommon flags are as follows:\n")
	args("-l --log", "show logs also on the standard output and errput\n")

	Log.Info("executed successfully")
}

func NewHelpProcess() *HelpProcess {
	Log.Info("created")
	return &HelpProcess{}
}
