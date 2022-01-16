package command

import "strconv"

type IntervalData struct {
	duration int
	rate     int
	filter   string
	script   string
}

func (i *IntervalData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Duration",
			DefaultValue: strconv.Itoa(i.duration),
			DataFunction: i.setDuration,
		},
		{
			Title:        "Rate",
			DefaultValue: strconv.Itoa(i.rate),
			DataFunction: i.setRate,
		},
		{
			Title:        "Filter",
			DefaultValue: i.filter,
			DataFunction: i.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: i.script,
			DataFunction: i.setScript,
		},
	}
}

func (i *IntervalData) setDuration(data string) {
	duration, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	i.duration = duration
}

func (i *IntervalData) setRate(data string) {
	rate, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	i.rate = rate
}

func (i *IntervalData) setFilter(data string) {
	i.filter = data
}

func (i *IntervalData) setScript(data string) {
	i.script = data
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
