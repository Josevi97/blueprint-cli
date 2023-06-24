package process

var (
	INIT      = "init"
	HELP      = "help"
	REGISTER  = "register"
	GENERATE  = "generate"
	LIST      = "list"
	BLUEPRINT = "blueprint"
)

type Process interface {
	RequiresToBeInitialized() bool

	// Every CLI process must include execute as the only one entrypoint
	Execute()
}
