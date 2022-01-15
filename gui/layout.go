package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

func addEnterCallback(g *gocui.Gui, vName string, cb func(index int) error) error {
	if err := g.SetKeybinding(vName, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, y := v.Cursor()
		return cb(y)
	}); err != nil {
		return err
	}
	return nil
}

func addEscCallback(g *gocui.Gui, vName string, cb func(index int) error) error {
	if err := g.SetKeybinding(vName, gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, y := v.Cursor()
		return cb(y)
	}); err != nil {
		return err
	}
	return nil
}

func addSelectKeyBindings(g *gocui.Gui, vName string) error {
	if err := g.SetKeybinding(vName, gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return selectItem(g, v, func(y, total int) int {
			return (y + total - 1) % total
		})
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(vName, gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return selectItem(g, v, func(y, total int) int {
			return (y + 1) % total
		})
	}); err != nil {
		return err
	}
	return nil
}

func setKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}

	if err := addSelectKeyBindings(g, "Choose"); err != nil {
		return err
	}
	if err := addEnterCallback(g, "Choose", func(index int) error {
		if guiData_.chooserData_.enterCb != nil {
			return guiData_.chooserData_.enterCb(index)
		}
		return nil
	}); err != nil {
		return err
	}

	if err := addEscCallback(g, "Choose", func(index int) error {
		if guiData_.chooserData_.escCb != nil {
			return guiData_.chooserData_.escCb(index)
		}
		return nil
	}); err != nil {
		return err
	}
	backCb := func(index int) error { backState(g); return nil }
	if err := addEnterCallback(g, "Tips", backCb); err != nil {
		return err
	}
	if err := addEscCallback(g, "Tips", backCb); err != nil {
		return err
	}

	if err := g.SetKeybinding("Choose", 'a', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if guiData_.state == stateTracerList {
			enterState(stateModuleList, g)
		}
		return nil
	}); err != nil {
		return err
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
