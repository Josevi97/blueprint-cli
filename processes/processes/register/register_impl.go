package Processes

import (
	logging "github.com/josevi97/core/logging/models"
	process "github.com/josevi97/processes/model"
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
