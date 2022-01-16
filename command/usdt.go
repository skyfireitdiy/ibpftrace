package command

type UsdtData struct {
	filepath  string
	namespace string
	name      string
	filter    string
	script    string
}

func (u *UsdtData) GenScript() string {
	return u.Script() + " { " + u.script + " }"
}

func (u *UsdtData) String() string {
	filterStr := ""
	if u.filter != "" {
		filterStr = " /" + u.filter + "/ "
	}
	namespace := ""
	if u.namespace != "" {
		namespace = u.namespace + ":"
	}
	return "usdt:" + u.filepath + ":" + namespace + u.name + filterStr
}

func (u *UsdtData) Script() string {
	return u.script
}
