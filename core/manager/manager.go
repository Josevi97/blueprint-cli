package manager

type Manager interface {
	// Method used to initialize the manager
	Init(path string)
}
