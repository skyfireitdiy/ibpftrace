package command

type KretprobeData struct {
	function string
	filter   string
	script   string
}

func (k *KretprobeData) GenScript() string {
	return k.String() + " { " + k.script + " }"
}

func (k *KretprobeData) String() string {
	filterStr := ""
	if k.filter != "" {
		filterStr = " /" + k.filter + "/ "
	}
	return "kretprobe:" + k.function + filterStr
}

func (k *KretprobeData) Script() string {
	return k.script
}
