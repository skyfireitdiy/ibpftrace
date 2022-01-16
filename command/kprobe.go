package command

import "strconv"

type KprobeData struct {
	function string
	offset   int
	filter   string
	script   string
}

func (k *KprobeData) GenScript() string {
	return k.String() + " { " + k.script + " }"
}

func (k *KprobeData) String() string {
	filterStr := ""
	if k.filter != "" {
		filterStr = " /" + k.filter + "/ "
	}
	return "kprobe:" + k.function + ":" + strconv.Itoa(k.offset) + filterStr
}

func (k *KprobeData) Script() string {
	return k.script
}
