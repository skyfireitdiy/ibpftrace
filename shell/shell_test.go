package shell

import "testing"

func TestRunShellCommand_NormalCommand(test *testing.T) {
	test.Log("TestRunShellCommand")
	stdout, stderr, err := RunShellCommand("echo Hello World")
	if err != nil {
		test.Error(err)
	}
	test.Log("stdout:", stdout)
	test.Log("stderr:", stderr)
	if stdout != "Hello World\n" {
		test.Error("stdout != Hello World\n")
	}
}

func TestRunShellCommand_ErrorCommand(test *testing.T) {
	test.Log("TestRunShellCommand")
	stdout, stderr, err := RunShellCommand("echo Hello World && exit 1")
	if err == nil {
		test.Error("err == nil")
	}
	test.Log("stdout:", stdout)
	test.Log("stderr:", stderr)
	if stdout != "Hello World\n" {
		test.Error("stdout != Hello World\n")
	}
}
