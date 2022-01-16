package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func updateData(g *gocui.Gui) {
	guiData_.viewConfig_[guiData_.state].updateFunc(g)
	updateStatusLayout(g)
}

func refreshUI(g *gocui.Gui) error {
	err := guiData_.viewConfig_[guiData_.state].refreshFunc(g)
	if err != nil {
		return err
	}
	return refreshStatusLayout(g)
}

func saveTracerListViewData() interface{} {
	return guiData_.tracerListData_
}

func saveModuleListViewData() interface{} {
	return guiData_.moduleListData_
}

func saveTipsViewData() interface{} {
	return guiData_.tipsData_
}

func recoverTracerListViewData(info stateInfo) {
	guiData_.tracerListData_ = info.data.(tracerListData)
}

func recoverModuleListViewData(info stateInfo) {
	guiData_.moduleListData_ = info.data.(moduleListData)
}

func recoverTipsViewData(info stateInfo) {
	guiData_.tipsData_ = info.data.(tipsData)
}

func enterState(state string, g *gocui.Gui) {
	info := stateInfo{}
	info.data = guiData_.viewConfig_[guiData_.state].enterSaveFunc()
	info.state = guiData_.state
	info.view = g.CurrentView().Name()
	guiData_.stateStack = append(guiData_.stateStack, info)
	guiData_.state = state
	updateData(g)
}

func convertState(state string, g *gocui.Gui) {
	guiData_.state = state
	updateData(g)
}

func backState(g *gocui.Gui) {
	if len(guiData_.stateStack) == 0 {
		guiData_.state = stateTracerList
		return
	}
	info := guiData_.stateStack[len(guiData_.stateStack)-1]
	guiData_.stateStack = guiData_.stateStack[:len(guiData_.stateStack)-1]
	guiData_.viewConfig_[info.state].backRecoverFunc(info)
	guiData_.state = info.state
	updateData(g)
}

func StatusString(g *gocui.Gui) string {
	ret := guiData_.state
	vName := g.CurrentView().Name()
	for _, hk := range guiData_.hotkey_ {
		if (hk.state == guiData_.state || hk.state == "") && (hk.view == vName || hk.view == "") {
			if hk.keyDisplayStr != "" {
				ret += " " + hk.keyDisplayStr + ":" + hk.desc
			} else {
				ret += fmt.Sprintf(" %c:%s", hk.key, hk.desc)
			}
			ret += "    "
		}
	}
	return ret
}
