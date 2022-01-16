package command

import "strconv"

type KprobeData struct {
	function string
	offset   int
	filter   string
	script   string
}

func (k *KprobeData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Function",
			DefaultValue: k.function,
			DataFunction: k.setFunction,
		},
		{
			Title:        "Offset",
			DefaultValue: strconv.Itoa(k.offset),
			DataFunction: k.setOffset,
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

func (k *KprobeData) setFunction(data string) {
	k.function = data
}

func (k *KprobeData) setOffset(data string) {
	offset, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	k.offset = offset
}

func (k *KprobeData) setFilter(data string) {
	k.filter = data
}

func (k *KprobeData) setScript(data string) {
	k.script = data
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
