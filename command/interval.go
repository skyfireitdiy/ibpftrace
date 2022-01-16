package command

import "strconv"

type IntervalData struct {
	duration int
	rate     int
	filter   string
	script   string
}

func (i *IntervalData) GenScript() string {
	return i.String() + " { " + i.script + " }"
}

func (i *IntervalData) String() string {
	filterStr := ""
	if i.filter != "" {
		filterStr = " /" + i.filter + "/ "
	}
	return "interval:" + strconv.Itoa(i.duration) + ":" + strconv.Itoa(i.rate) + filterStr
}

func (i *IntervalData) Script() string {
	return i.script
}
