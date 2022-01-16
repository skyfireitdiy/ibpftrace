package command

import "strconv"

type SoftwareData struct {
	eventName string
	count     int
	filter    string
	script    string
}

func (s *SoftwareData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Event Name",
			DefaultValue: s.eventName,
			DataFunction: s.setEventName,
		},
		{
			Title:        "Count",
			DefaultValue: strconv.Itoa(s.count),
			DataFunction: s.setCount,
		},
		{
			Title:        "Filter",
			DefaultValue: s.filter,
			DataFunction: s.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: s.script,
			DataFunction: s.setScript,
		},
	}
}

func (s *SoftwareData) setScript(data string) {
	s.script = data
}

func (s *SoftwareData) setEventName(data string) {
	s.eventName = data
}

func (s *SoftwareData) setCount(data string) {
	count, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	s.count = count
}

func (s *SoftwareData) setFilter(data string) {
	s.filter = data
}

func (s *SoftwareData) GenScript() string {
	return s.String() + " { " + s.script + " }"
}

func (s *SoftwareData) String() string {
	filterStr := ""
	if s.filter != "" {
		filterStr = " /" + s.filter + "/ "
	}
	return "software:" + s.eventName + ":" + strconv.Itoa(s.count) + filterStr
}

func (s *SoftwareData) Script() string {
	return s.script
}
