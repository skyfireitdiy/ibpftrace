package utils

import "testing"

func TestBpftraceExist(test *testing.T) {
	test.Log("TestBpftraceExist")
	if bpfTraceExist() {
		test.Log("bpfTraceExist() == true")
	} else {
		test.Error("bpfTraceExist() == false")
	}
}

func TestIsRoot(test *testing.T) {
	test.Log("TestIsRoot")
	if isRoot() {
		test.Log("isRoot() == true")
	} else {
		test.Error("isRoot() == false")
	}
}
