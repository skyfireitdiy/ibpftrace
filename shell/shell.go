package shell

import (
	"bytes"
	"os/exec"
)

func RunShellCommand(command string) (string, string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return out.String(), stderr.String(), err
}
