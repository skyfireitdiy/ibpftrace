package command

import "strconv"

type ProfileData struct {
	duration int
	rate     int
	filter   string
	script   string
}

func (p *ProfileData) GenScript() string {
	return p.String() + " { " + p.script + " }"
}

func (p *ProfileData) String() string {
	filterStr := ""
	if p.filter != "" {
		filterStr = " /" + p.filter + "/ "
	}
	return "profile:" + strconv.Itoa(p.duration) + ":" + strconv.Itoa(p.rate) + filterStr
}

func (p *ProfileData) Script() string {
	return p.script
}
