package gui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

func syncModuleListInfo(g *gocui.Gui) {
	guiData_.chooserData_.escCb = guiData_.moduleListData_.escCb
	guiData_.chooserData_.enterCb = guiData_.moduleListData_.enterCb
	guiData_.chooserData_.selChangedCb = nil
	guiData_.chooserData_.status = guiData_.moduleListData_.status

	guiData_.chooserData_.title = stateTracerList
	guiData_.chooserData_.data = strings.Join(guiData_.moduleListData_.moduleList, "\n")
	guiData_.chooserData_.y = guiData_.moduleListData_.y

	maxX, maxY := g.Size()
	guiData_.chooserData_.x0 = 0
	guiData_.chooserData_.y0 = 0
	guiData_.chooserData_.x1 = maxX - 1
	guiData_.chooserData_.y1 = maxY - 5

	guiData_.statusData_.content = guiData_.moduleListData_.status
}

func refreshModuleListLayout(g *gocui.Gui) error {
	guiData_.moduleListData_.enterCb = func(index int) error {
		guiData_.moduleListData_.y = index
		setTips("Tips", fmt.Sprintf("Module %s not implemented yet.", guiData_.moduleListData_.moduleList[index]))
		convertTips(g)
		return nil
	}
	guiData_.moduleListData_.escCb = func(index int) error {
		return nil
	}
	guiData_.moduleListData_.status = "Choose Module"
	syncModuleListInfo(g)
	return chooseLayout(g)
}

func updateModuleListLayout(g *gocui.Gui) {
	syncModuleListInfo(g)
	updateChooseLayout(g)
}
