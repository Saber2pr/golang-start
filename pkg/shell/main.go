package shell

import (
	"bytes"
	"os/exec"
)

func ExecCommand(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}

func Ls(path string) (string, error) {
	return ExecCommand("ls " + path)
}
