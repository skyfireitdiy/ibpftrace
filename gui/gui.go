package gui

import (
	"github.com/jroimartin/gocui"
)

func updateData(g *gocui.Gui) {
	switch guiData_.state {
	case stateTracerList:
		updateTracerListLayout(g)
	case stateModuleList:
		updateModuleListLayout(g)
	case stateTips:
		updateTipsLayout(g)
	}
	updateStatusLayout(g)
}

func refreshUI(g *gocui.Gui) error {
	var err error
	switch guiData_.state {
	case stateTracerList:
		err = refreshTracerListLayout(g)
	case stateModuleList:
		err = refreshModuleListLayout(g)
	case stateTips:
		err = refreshTipsLayout(g)
	}
	if err != nil {
		return err
	}
	return refreshStatusLayout(g)
}

func enterState(state string, g *gocui.Gui) {
	info := stateInfo{}
	switch guiData_.state {
	case stateTracerList:
		info.data = guiData_.tracerListData_
	case stateModuleList:
		info.data = guiData_.moduleListData_
	case stateTips:
		info.data = guiData_.tipsData_
	default:
	}
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
	switch info.state {
	case stateTracerList:
		guiData_.tracerListData_ = info.data.(tracerListData)
	case stateModuleList:
		guiData_.moduleListData_ = info.data.(moduleListData)
	case stateTips:
		guiData_.tipsData_ = info.data.(tipsData)
	}
	guiData_.state = info.state
	updateData(g)
}
