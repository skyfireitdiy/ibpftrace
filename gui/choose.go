package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func selectItem(g *gocui.Gui, v *gocui.View, convert func(y, total int) int) error {
	_, y := v.Cursor()
	total := len(v.BufferLines()) - 1
	if total == 0 {
		return nil
	}
	y = convert(y, total)
	v.SetCursor(0, y)
	if guiData_.chooserData_.selChangedCb != nil {
		guiData_.chooserData_.selChangedCb(y)
	}
	return nil
}

func chooseLayout(g *gocui.Gui) error {
	if v, err := g.SetView("Choose", guiData_.chooserData_.x0, guiData_.chooserData_.y0, guiData_.chooserData_.x1, guiData_.chooserData_.y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
		v.FgColor = gocui.ColorGreen
		v.BgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorWhite
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true

		updateChooseContent(g)
	}
	return nil
}

func updateChooseContent(g *gocui.Gui) error {
	v, err := g.View("Choose")
	if err != nil {
		return err
	}
	v.Title = guiData_.chooserData_.title
	v.Clear()
	fmt.Fprintln(v, guiData_.chooserData_.data)
	v.SetCursor(0, guiData_.chooserData_.y)
	g.SetCurrentView("Choose")
	g.SetViewOnTop("Choose")
	return nil
}

func updateChooseLayout(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		return updateChooseContent(g)
	})
}
