package gui

import (
	"github.com/jroimartin/gocui"
	"go.skyfire.com/command"
)

const (
	stateTracerList = "Module List"
	stateModuleList = "Choose Module"
	stateTips       = "Tips"
)

type viewConfig struct {
	updateFunc      func(g *gocui.Gui)
	refreshFunc     func(g *gocui.Gui) error
	enterSaveFunc   func() interface{}
	backRecoverFunc func(info stateInfo)
}

var guiData_ guiData

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

	viewConfig_ map[string]viewConfig

	moduleListData_ moduleListData
	tipsData_       tipsData
	tracerListData_ tracerListData
	statusData_     statusData

	stateStack []stateInfo
}

func InitGuiData() {
	guiData_ = guiData{
		state: stateTracerList,
		moduleListData_: moduleListData{
			moduleList: command.ModuleList(),
		},
		tipsData_: tipsData{},
		viewConfig_: map[string]viewConfig{
			stateTracerList: {
				updateFunc:      updateTracerListLayout,
				refreshFunc:     refreshTracerListLayout,
				enterSaveFunc:   saveTracerListViewData,
				backRecoverFunc: recoverTracerListViewData,
			},
			stateModuleList: {
				updateFunc:      updateModuleListLayout,
				refreshFunc:     refreshModuleListLayout,
				enterSaveFunc:   saveModuleListViewData,
				backRecoverFunc: recoverModuleListViewData,
			},
			stateTips: {
				updateFunc:      updateTipsLayout,
				refreshFunc:     refreshTipsLayout,
				enterSaveFunc:   saveTipsViewData,
				backRecoverFunc: recoverTipsViewData,
			},
		},
	}
}
