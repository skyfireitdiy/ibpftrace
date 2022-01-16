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
}

type tipsData struct {
	title   string
	content string
}

type tracerListData struct {
	enterCb      func(index int) error
	selChangedCb func(index int) error
	y            int
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
}

type hotkey struct {
	view          string
	desc          string
	key           interface{}
	keyDisplayStr string
	state         string
	mod           gocui.Modifier
	handler       func(g *gocui.Gui, v *gocui.View) error
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
	hotkey_    []hotkey
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
		hotkey_: []hotkey{
			{
				key:           gocui.KeyCtrlC,
				desc:          "Quit",
				keyDisplayStr: "Ctrl+C",
				mod:           gocui.ModNone,
				handler:       quit,
			},
			{
				key:     'q',
				desc:    "Quit",
				mod:     gocui.ModNone,
				handler: quit,
			},
			{
				view:          "Choose",
				desc:          "Down",
				keyDisplayStr: "↓",
				key:           gocui.KeyArrowDown,
				mod:           gocui.ModNone,
				handler: func(g *gocui.Gui, v *gocui.View) error {
					return selectItem(g, v, func(y, total int) int {
						return (y + 1) % total
					})
				},
			},
			{
				view:          "Choose",
				desc:          "Up",
				keyDisplayStr: "↑",
				key:           gocui.KeyArrowUp,
				mod:           gocui.ModNone,
				handler: func(g *gocui.Gui, v *gocui.View) error {
					return selectItem(g, v, func(y, total int) int {
						return (y + total - 1) % total
					})
				},
			},
			{
				view:          "Choose",
				desc:          "Choose",
				key:           gocui.KeyEnter,
				mod:           gocui.ModNone,
				keyDisplayStr: "Enter",
				handler: func(g *gocui.Gui, v *gocui.View) error {
					_, index := v.Cursor()
					if guiData_.chooserData_.enterCb != nil {
						return guiData_.chooserData_.enterCb(index)
					}
					return nil
				},
			},
			{
				view:          "Choose",
				desc:          "Back",
				key:           gocui.KeyEsc,
				mod:           gocui.ModNone,
				keyDisplayStr: "Esc",
				handler: func(g *gocui.Gui, v *gocui.View) error {
					_, index := v.Cursor()
					if guiData_.chooserData_.escCb != nil {
						return guiData_.chooserData_.escCb(index)
					}
					return nil
				},
			},
			{
				view:          "Tips",
				desc:          "Back",
				key:           gocui.KeyEnter,
				mod:           gocui.ModNone,
				keyDisplayStr: "Enter",
				handler: func(g *gocui.Gui, v *gocui.View) error {
					backState(g)
					return nil
				},
			},
			{
				view:          "Tips",
				desc:          "Back",
				key:           gocui.KeyEsc,
				mod:           gocui.ModNone,
				keyDisplayStr: "Esc",
				handler: func(g *gocui.Gui, v *gocui.View) error {
					backState(g)
					return nil
				},
			},
			{
				view:  "Choose",
				desc:  "Add Tracer",
				key:   'a',
				mod:   gocui.ModNone,
				state: stateTracerList,
				handler: func(g *gocui.Gui, v *gocui.View) error {
					enterState(stateModuleList, g)
					return nil
				},
			},
		},
	}
}
