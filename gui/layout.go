package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

func makeHandler(hk hotkey) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if guiData_.state == hk.state || hk.state == "" {
			return hk.handler(g, v)
		}
		return nil
	}
}

func setKeyBindings(g *gocui.Gui) error {
	for _, hk := range guiData_.hotkey_ {
		if err := g.SetKeybinding(hk.view, hk.key, gocui.ModNone, makeHandler(hk)); err != nil {
			return err
		}
	}
	return nil
}

func GuiLoop() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(refreshUI)

	if err := setKeyBindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
