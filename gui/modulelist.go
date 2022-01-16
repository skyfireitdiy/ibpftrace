package gui

import (
	"strings"

	"github.com/jroimartin/gocui"
	"go.skyfire.com/command"
)

func syncModuleListInfo(g *gocui.Gui) {
	guiData_.chooserData_.escCb = guiData_.moduleListData_.escCb
	guiData_.chooserData_.enterCb = guiData_.moduleListData_.enterCb
	guiData_.chooserData_.selChangedCb = nil

	guiData_.chooserData_.title = stateTracerList
	guiData_.chooserData_.data = strings.Join(guiData_.moduleListData_.moduleList, "\n")
	guiData_.chooserData_.y = guiData_.moduleListData_.y

	maxX, maxY := g.Size()
	guiData_.chooserData_.x0 = 0
	guiData_.chooserData_.y0 = 0
	guiData_.chooserData_.x1 = maxX - 1
	guiData_.chooserData_.y1 = maxY - 5

}

func refreshModuleListLayout(g *gocui.Gui) error {
	guiData_.moduleListData_.enterCb = func(index int) error {
		guiData_.moduleListData_.y = index
		tracer := command.CreateTracer(index)
		command.AddTracer(tracer)
		backState(g)
		return nil
	}
	guiData_.moduleListData_.escCb = func(index int) error {
		return nil
	}
	syncModuleListInfo(g)
	return refreshChooseLayout(g)
}

func updateModuleListLayout(g *gocui.Gui) {
	syncModuleListInfo(g)
	updateChooseLayout(g)
}
