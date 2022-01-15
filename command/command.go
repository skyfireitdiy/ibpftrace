package command

type tracer interface {
	GenScript() string
	Script() string
	String() string
}

var tracerList = []tracer{}

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

func GetChooseList() []string {
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
