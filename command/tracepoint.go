package command

type TracepointData struct {
	subsys string
	event  string
	filter string
	script string
}

func (t *TracepointData) GenScript() string {
	return t.String() + " { " + t.script + " }"
}

func (t *TracepointData) String() string {
	filterStr := ""
	if t.filter != "" {
		filterStr = " /" + t.filter + "/ "
	}
	return "tracepoint:" + t.subsys + ":" + t.event + filterStr
}

func (t *TracepointData) Script() string {
	return t.script
}
