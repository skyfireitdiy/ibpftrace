package gui

import (
	"strings"

	"github.com/jroimartin/gocui"
	"go.skyfire.com/command"
)

func syncTracerListInfo(g *gocui.Gui) {
	guiData_.chooserData_.escCb = nil
	guiData_.chooserData_.enterCb = guiData_.tracerListData_.enterCb
	guiData_.chooserData_.selChangedCb = guiData_.tracerListData_.selChangedCb

	guiData_.chooserData_.title = stateTracerList
	guiData_.chooserData_.data = strings.Join(command.TracerList(), "\n")
	guiData_.chooserData_.y = guiData_.tracerListData_.y

	maxX, maxY := g.Size()
	guiData_.chooserData_.x0 = 0
	guiData_.chooserData_.y0 = 0
	guiData_.chooserData_.x1 = maxX/2 - 1
	guiData_.chooserData_.y1 = maxY - 5

}

func refreshTracerListLayout(g *gocui.Gui) error {
	guiData_.tracerListData_.enterCb = func(index int) error {
		guiData_.tracerListData_.y = index
		return nil
	}
	guiData_.tracerListData_.selChangedCb = func(index int) error {
		return nil
	}
	syncTracerListInfo(g)
	if err := chooseLayout(g); err != nil {
		return err
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("Script", maxX/2, 0, maxX-1, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
		v.Title = "Script"
		updateScriptContent(g)
	}
	return nil
}

func updateScriptContent(g *gocui.Gui) {
	v, err := g.View("Script")
	if err != nil {
		return
	}
	v.Clear()
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	v.Write([]byte(command.Script(guiData_.tracerListData_.y)))
	g.SetViewOnTop("Script")
}

func updateTracerListLayout(g *gocui.Gui) {
	syncTracerListInfo(g)
	updateScriptContent(g)
	updateChooseLayout(g)
}
