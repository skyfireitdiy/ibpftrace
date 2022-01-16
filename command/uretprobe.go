package command

type UretprobeData struct {
	filepath string
	function string
	filter   string
	script   string
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
