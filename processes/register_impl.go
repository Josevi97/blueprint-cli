package processes

type RegisterProcess struct{}

func (registerProcess *RegisterProcess) RequiresToBeInitialized() bool {
	return true
}

func (registerProcess *RegisterProcess) Execute() {
	Log.Info("executed successfully")
}

func NewRegisterProcess() Process {
	return &RegisterProcess{}
}
