package command

import "strconv"

type ProfileData struct {
	duration int
	rate     int
	filter   string
	script   string
}

func (p *ProfileData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Duration",
			DefaultValue: strconv.Itoa(p.duration),
			DataFunction: p.setDuration,
		},
		{
			Title:        "Rate",
			DefaultValue: strconv.Itoa(p.rate),
			DataFunction: p.setRate,
		},
		{
			Title:        "Filter",
			DefaultValue: p.filter,
			DataFunction: p.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: p.script,
			DataFunction: p.setScript,
		},
	}
}

func (p *ProfileData) setDuration(data string) {
	duration, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	p.duration = duration
}

func (p *ProfileData) setRate(data string) {
	rate, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	p.rate = rate
}

func (p *ProfileData) setFilter(data string) {
	p.filter = data
}

func (p *ProfileData) setScript(data string) {
	p.script = data
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
