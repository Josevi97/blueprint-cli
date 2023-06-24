package CommandUtils

import (
	"os/exec"
	"strings"
)

func Pwd() string {
	output, err := exec.Command("pwd").Output()

	if err != nil {
		return ""
	}

	return strings.ReplaceAll(string(output), "\n", "")
}
