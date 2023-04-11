package os_cmds

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	Windows = "windows"
	Darwin  = "darwin"
	Linux   = "linux"
	Android = "android"
)

func OpenPdf(filename string) error {
	var c *exec.Cmd

	switch runtime.GOOS {
	case Linux:
		c = exec.Command("xdg-open", filename)
	case Windows:
		c = exec.Command("powershell", "-command", fmt.Sprintf("start '%s'", filename))
	default:
		return fmt.Errorf("unsupported OS")
	}

	return c.Run()
}
