package process_manager

type Process interface {
	// Every cli process must include execute as the only one entrypoint
	Execute()
}
