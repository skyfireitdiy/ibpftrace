package command

type UprobeData struct {
	filepath string
	function string
	offset   int
	filter   string
	script   string
}

func (u *UprobeData) GenScript() string {
	return u.Script() + " { " + u.script + " }"
}

func (u *UprobeData) String() string {
	filterStr := ""
	if u.filter != "" {
		filterStr = " /" + u.filter + "/ "
	}
	return "uretprobe:" + u.filepath + ":" + u.function + filterStr
}

func (u *UprobeData) Script() string {
	return u.script
}
