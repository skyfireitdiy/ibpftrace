package command

type tracer interface {
	GenScript() string
	Script() string
	String() string
}

type module struct {
	name          string
	desc          string
	tracerCreator func() tracer
}

var tracerList = []tracer{}
var moduleList = []module{
	{
		name: "kprobe",
		desc: "kprobe - kernel function start",
		tracerCreator: func() tracer {
			return &KprobeData{}
		},
	},
	{
		name: "kretprobe",
		desc: "kretprobe - kernel function return",
		tracerCreator: func() tracer {
			return &KretprobeData{}
		},
	},
	{
		name: "uprobe",
		desc: "uprobe - user-level function start",
		tracerCreator: func() tracer {
			return &UprobeData{}
		},
	},
	{
		name: "tracepoint",
		desc: "tracepoint - kernel static tracepoints",
		tracerCreator: func() tracer {
			return &TracepointData{}
		},
	},
	{
		name: "uretprobe",
		desc: "uretprobe - user-level function return",
		tracerCreator: func() tracer {
			return &UretprobeData{}
		},
	},
	{
		name: "usdt",
		desc: "usdt - user-level static tracepoints",
		tracerCreator: func() tracer {
			return &UsdtData{}
		},
	},
	{
		name: "profile",
		desc: "profile - timed sampling",
		tracerCreator: func() tracer {
			return &ProfileData{}
		},
	},
	{
		name: "interval",
		desc: "interval - timed output",
		tracerCreator: func() tracer {
			return &IntervalData{}
		},
	},
	{
		name: "software",
		desc: "software - kernel software events",
		tracerCreator: func() tracer {
			return &SoftwareData{}
		},
	},
	{
		name: "hardware",
		desc: "hardware - processor-level events",
		tracerCreator: func() tracer {
			return &HardwareData{}
		},
	},
}

func AddTracer(module tracer) {
	tracerList = append(tracerList, module)
}

func RemoveTracer(index int) {
	if index >= len(tracerList) {
		return
	}
	tracerList = append(tracerList[:index], tracerList[index+1:]...)
}

func ReplaceTracer(index int, module tracer) {
	if index >= len(tracerList) {
		return
	}
	tracerList[index] = module
}

func TracerList() []string {
	ret := []string{}
	for _, module := range tracerList {
		ret = append(ret, module.String())
	}
	return ret
}

func Script(index int) string {
	if index >= len(tracerList) {
		return ""
	}
	return tracerList[index].Script()
}

func ModuleList() []string {
	ret := []string{}
	for _, module := range moduleList {
		ret = append(ret, module.desc)
	}
	return ret
}
