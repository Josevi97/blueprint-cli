package HelpProcess

import (
	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/process"
)

type HelpProcess struct{}

var Log = logging.NewLogging("HELP PROCESS")

func commandHelp(command string, msg string) {
	Log.Output("%s\t%s", command, msg)
}

func (helpProcess *HelpProcess) Execute() {
	Log.Output("Blueprint CLI for structuring your project defining conventions.")
	Log.Output("The basic CLI commands are as follows:")
	commandHelp(process.HELP, "\tdisplays the blueprint CLI commands available")
	commandHelp(process.INIT, "\tuse this command to initialize blueprint CLI")
	commandHelp(process.REGISTER, "this command is used to register a new subcommand")
	commandHelp(process.GENERATE, "generates a structure depending on the subcommands registered")
	commandHelp(process.BLUEPRINT, "opens GUI blueprint")

	Log.Info("executed successfully")
}

func NewHelpProcess() *HelpProcess {
	Log.Info("created")
	return &HelpProcess{}
}
