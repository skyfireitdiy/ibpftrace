package gui

const (
	stateTracerList = "Module List"
	stateModuleList = "Choose Module"
	stateTips       = "Tips"
)

var guiData_ = guiData{
	state: stateTracerList,
	moduleListData_: moduleListData{
		moduleList: []string{
			"kprobe - kernel function start",
			"kretprobe - kernel function return",
			"uprobe - user-level function start",
			"uretprobe - user-level function return",
			"tracepoint - kernel static tracepoints",
			"usdt - user-level static tracepoints",
			"profile - timed sampling",
			"interval - timed output",
			"software - kernel software events",
			"hardware - processor-level events",
		},
	},
	tipsData_: tipsData{},
}

type moduleListData struct {
	moduleList []string
	enterCb    func(index int) error
	escCb      func(index int) error
	y          int
	status     string
}

type tipsData struct {
	title   string
	content string
	status  string
}

type tracerListData struct {
	enterCb      func(index int) error
	selChangedCb func(index int) error
	y            int
	status       string
}

type statusData struct {
	content string
}

type stateInfo struct {
	state string
	view  string
	data  interface{}
}

type chooserData struct {
	selChangedCb   func(index int) error
	escCb          func(index int) error
	enterCb        func(index int) error
	title          string
	data           interface{}
	x0, y0, x1, y1 int
	y              int
	status         string
}

type guiData struct {
	state        string
	chooserData_ chooserData

	moduleListData_ moduleListData
	tipsData_       tipsData
	tracerListData_ tracerListData
	statusData_     statusData

	stateStack []stateInfo
}
