package command

import "strconv"

type HardwareData struct {
	eventName string
	count     int
	filter    string
	script    string
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
