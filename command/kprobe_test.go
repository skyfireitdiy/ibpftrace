package command

import "testing"

func TestKprobeGenString(test *testing.T) {
	data := KprobeData{
		function: "_do_fork",
		offset:   0,
		filter:   "",
		script:   `printf("%d\n", tid);`,
	}

	if data.GenScript() != `kprobe:_do_fork:0 { printf("%d\n", tid); }` {
		test.Error("Kprobe script is not correct", data.GenScript())
	}
}
