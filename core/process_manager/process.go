package ProcessManager

type Process interface {
	// Every cli process must include execute as the only one entrypoint
	Execute()
}
