package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func refreshTipsLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Tips", 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = guiData_.tipsData_.title
		v.Wrap = true
		v.Autoscroll = true
		updateTipsLayout(g)
	}
	return nil
}

func setTips(title, content string) {
	guiData_.tipsData_.title = title
	guiData_.tipsData_.content = content
}

func enterTips(g *gocui.Gui) {
	enterState(stateTips, g)
}

func convertTips(g *gocui.Gui) {
	convertState(stateTips, g)
}

func updateTipsLayout(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("Tips")
		if err != nil {
			return err
		}
		v.Title = guiData_.tipsData_.title
		v.Clear()
		fmt.Fprintln(v, guiData_.tipsData_.content)
		g.SetCurrentView("Tips")
		g.SetViewOnTop("Tips")
		return nil
	})
}
