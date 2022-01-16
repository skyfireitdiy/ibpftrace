package command

import "strconv"

type HardwareData struct {
	eventName string
	count     int
	filter    string
	script    string
}

func (h *HardwareData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Event Name",
			DefaultValue: h.eventName,
			DataFunction: h.setEventName,
		},
		{
			Title:        "Count",
			DefaultValue: strconv.Itoa(h.count),
			DataFunction: h.setCount,
		},
		{
			Title:        "Filter",
			DefaultValue: h.filter,
			DataFunction: h.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: h.script,
			DataFunction: h.setScript,
		},
	}
}

func (h *HardwareData) setScript(data string) {
	h.script = data
}

func (h *HardwareData) setEventName(data string) {
	h.eventName = data
}

func (h *HardwareData) setCount(data string) {
	count, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	h.count = count
}

func (h *HardwareData) setFilter(data string) {
	h.filter = data
}

func (h *HardwareData) GenScript() string {
	return h.String() + " { " + h.script + " }"
}

func (h *HardwareData) String() string {
	filterStr := ""
	if h.filter != "" {
		filterStr = " /" + h.filter + "/ "
	}
	return "hardware:" + h.eventName + ":" + strconv.Itoa(h.count) + filterStr
}

func (h *HardwareData) Script() string {
	return h.script
}
