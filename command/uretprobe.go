package command

type UretprobeData struct {
	filepath string
	function string
	filter   string
	script   string
}

func (u *UretprobeData) EditItems() []EditItem {
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

func (u *UretprobeData) setFilepath(data string) {
	u.filepath = data
}

func (u *UretprobeData) setFunction(data string) {
	u.function = data
}

func (u *UretprobeData) setFilter(data string) {
	u.filter = data
}

func (u *UretprobeData) setScript(data string) {
	u.script = data
}

func (u *UretprobeData) GenScript() string {
	return u.String() + " { " + u.script + " }"
}

func (u *UretprobeData) String() string {
	filterStr := ""
	if u.filter != "" {
		filterStr = " /" + u.filter + "/ "
	}
	return "uretprobe:" + u.filepath + ":" + u.function + filterStr
}

func (u *UretprobeData) Script() string {
	return u.script
}
