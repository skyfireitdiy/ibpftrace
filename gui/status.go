package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func refreshStatusLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Status", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		updateStatusContent(g)
	}
	return nil
}

func updateStatusContent(g *gocui.Gui) error {
	v, err := g.View("Status")
	if err != nil {
		return err
	}
	v.Clear()
	v.SetCursor(0, 0)
	fmt.Fprintln(v, StatusString(g))
	return nil
}

func updateStatusLayout(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		return updateStatusContent(g)
	})
}
