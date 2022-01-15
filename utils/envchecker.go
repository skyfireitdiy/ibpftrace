package utils

import (
	"log"
	"strings"

	"go.skyfire.com/shell"
)

func bpfTraceExist() bool {
	_, _, err := shell.RunShellCommand("which bpftrace")
	if err != nil {
		log.Println("bpftrace not found")
		return false
	}
	return true
}

func isRoot() bool {
	out, _, err := shell.RunShellCommand("id -u")
	if err != nil || strings.TrimSpace(out) != "0" {
		log.Printf("Not root, id -u: %s", out)
		log.Printf("this command requires root privilege")
		return false
	}
	return true
}

func EnvCheck() bool {
	return bpfTraceExist() && isRoot()
}
