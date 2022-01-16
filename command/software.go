package command

import "strconv"

type SoftwareData struct {
	eventName string
	count     int
	filter    string
	script    string
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
