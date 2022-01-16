package command

type KretprobeData struct {
	function string
	filter   string
	script   string
}

func (k *KretprobeData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Function",
			DefaultValue: k.function,
			DataFunction: k.setFunction,
		},
		{
			Title:        "Filter",
			DefaultValue: k.filter,
			DataFunction: k.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: k.script,
			DataFunction: k.setScript,
		},
	}
}

func (k *KretprobeData) setFunction(data string) {
	k.function = data
}

func (k *KretprobeData) setFilter(data string) {
	k.filter = data
}

func (k *KretprobeData) setScript(data string) {
	k.script = data
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
