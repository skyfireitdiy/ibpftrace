package command

type TracepointData struct {
	subsys string
	event  string
	filter string
	script string
}

func (t *TracepointData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Subsystem",
			DefaultValue: t.subsys,
			DataFunction: t.setSubsys,
		},
		{
			Title:        "Event",
			DefaultValue: t.event,
			DataFunction: t.setEvent,
		},
		{
			Title:        "Filter",
			DefaultValue: t.filter,
			DataFunction: t.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: t.script,
			DataFunction: t.setScript,
		},
	}
}

func (t *TracepointData) setSubsys(data string) {
	t.subsys = data
}

func (t *TracepointData) setEvent(data string) {
	t.event = data
}

func (t *TracepointData) setFilter(data string) {
	t.filter = data
}

func (t *TracepointData) setScript(data string) {
	t.script = data
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
