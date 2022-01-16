package command

import "strconv"

type UprobeData struct {
	filepath string
	function string
	offset   int
	filter   string
	script   string
}

func (u *UprobeData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Filepath",
			DefaultValue: u.filepath,
			DataFunction: u.setFilepath,
		},
		{
			Title:        "Function",
			DefaultValue: u.function,
			DataFunction: u.setFunction,
		},
		{
			Title:        "Offset",
			DefaultValue: strconv.Itoa(u.offset),
			DataFunction: u.setOffset,
		},
		{
			Title:        "Filter",
			DefaultValue: u.filter,
			DataFunction: u.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: u.script,
			DataFunction: u.setScript,
		},
	}
}

func (u *UprobeData) setFilepath(data string) {
	u.filepath = data
}

func (u *UprobeData) setFunction(data string) {
	u.function = data
}

func (u *UprobeData) setOffset(data string) {
	offset, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	u.offset = offset
}

func (u *UprobeData) setFilter(data string) {
	u.filter = data
}

func (u *UprobeData) setScript(data string) {
	u.script = data
}

func (u *UprobeData) GenScript() string {
	return u.Script() + " { " + u.script + " }"
}

func (u *UprobeData) String() string {
	filterStr := ""
	if u.filter != "" {
		filterStr = " /" + u.filter + "/ "
	}
	return "uretprobe:" + u.filepath + ":" + u.function + "+" + strconv.Itoa(u.offset) + filterStr
}

func (u *UprobeData) Script() string {
	return u.script
}
