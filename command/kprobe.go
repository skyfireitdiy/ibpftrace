package command

import "strconv"

type KprobeData struct {
	function string
	offset   int
	filter   string
	script   string
}

func (k *KprobeData) GenScript() string {
	filterStr := ""
	if k.filter != "" {
		filterStr = " /" + k.filter + "/ "
	}
	return "kprobe:" + k.function + ":" + strconv.Itoa(k.offset) + filterStr + " { " + k.script + " }"
}

func (k *KprobeData) String() string {
	return k.function + ":" + strconv.Itoa(k.offset) + ":" + k.filter
}

func (k *KprobeData) Script() string {
	return k.script
}
