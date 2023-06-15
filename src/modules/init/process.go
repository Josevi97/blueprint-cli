package init_process

type InitProcess struct{}

func (init *InitProcess) Execute() {

}

func NewInitProcess() *InitProcess {
	return &InitProcess{}
}
