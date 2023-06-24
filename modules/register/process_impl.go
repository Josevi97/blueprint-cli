package register

import (
	"github.com/josevi97/core/logging"
	"github.com/josevi97/core/process"
)

type RegisterProcess struct{}

var Log = logging.NewLogging("REGISTER PROCESS")

func (registerProcess *RegisterProcess) RequiresToBeInitialized() bool {
	return true
}

func (registerProcess *RegisterProcess) Execute() {
	Log.Info("executed successfully")
}

func NewRegisterProcess() process.Process {
	return &RegisterProcess{}
}
