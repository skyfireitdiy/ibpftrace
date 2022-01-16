package command

type UsdtData struct {
	filepath  string
	namespace string
	name      string
	filter    string
	script    string
}

func (u *UsdtData) EditItems() []EditItem {
	return []EditItem{
		{
			Title:        "Filepath",
			DefaultValue: u.filepath,
			DataFunction: u.setFilepath,
		},
		{
			Title:        "Namespace",
			DefaultValue: u.namespace,
			DataFunction: u.setNamespace,
		},
		{
			Title:        "Name",
			DefaultValue: u.name,
			DataFunction: u.setName,
		},
		{
			Title:        "Filter",
			DefaultValue: u.filter,
			DataFunction: u.setFilter,
		},
		{
			Title:        "Script",
			DefaultValue: u.script,
			DataFunction: u.setScript,
		},
	}
}

func (u *UsdtData) setFilepath(data string) {
	u.filepath = data
}

func (u *UsdtData) setNamespace(data string) {
	u.namespace = data
}

func (u *UsdtData) setName(data string) {
	u.name = data
}

func (u *UsdtData) setFilter(data string) {
	u.filter = data
}

func (u *UsdtData) setScript(data string) {
	u.script = data
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
