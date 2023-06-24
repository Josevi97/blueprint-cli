package HelpProcess

import (
	logging "github.com/josevi97/core/logging/models"
	process "github.com/josevi97/processes/model"
)

type HelpProcess struct{}

var Log = logging.NewLogging("HELP PROCESS")

func (helpProcess *HelpProcess) RequiresToBeInitialized() bool {
	return false
}

func args(command string, msg string) {
	Log.Output("%s\t\t%s\n", command, msg)
}

func outHeader() {
	Log.Output("Blueprint CLI for structuring your project defining conventions.")
	Log.Output("The basic CLI commands are as follows:\n")
}

func outHelp() {
	args(process.HELP, "\tdisplays the blueprint CLI commands available")
}

func outInit() {
	args(process.INIT, "\tuse this command to initialize blueprint CLI")
}

func outRegister() {
	args(process.REGISTER, "registers a row in the database required for diffrent things")
	args("  args", "\t[directory name] [command name]")

	args("  -y", "\tthis registers a command with everything as default. If not, an assistant will help you to register it")
	args("  --command", "set the registration for a command. This is will used by default")
	args("  -p", "\tpath to generate the registered action. If flag is not used, it will be required when using generate command\n")
	args("  -t", "\tfile template to be used. This can be used multiple times")
	args("  -g", "\tsubgenerate command. It must be registered firstly to work correctly. This can be used multiple times")
	args("  --template", "set the registration for a template")
}

func outGenerate() {
	args(process.GENERATE, "generates a block item depending on the blocks registered")
}

func outList() {
	args(process.LIST, "\tlist every command block generated")
}

func outBlueprint() {
	args(process.BLUEPRINT, "opens GUI blueprint")
}

func outBody() {
	outHelp()

	outInit()

	outRegister()

	outGenerate()

	outList()

	outBlueprint()
}

func outFooter() {
	Log.Output("\nCommon flags are as follows:\n")
	args("-l", "show logs also on the standard output and errput\n")
}

func (helpProcess *HelpProcess) Execute() {
	outHeader()
	outBody()
	outFooter()

	Log.Info("executed successfully")
}

func NewHelpProcess() *HelpProcess {
	Log.Info("created")
	return &HelpProcess{}
}
