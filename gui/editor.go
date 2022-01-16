package gui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

func refreshEditorLayout(g *gocui.Gui) error {
	if guiData_.editorData_.curr >= len(guiData_.editorData_.items) {
		return fmt.Errorf("index error: %d", guiData_.editorData_.curr)
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("Editor", 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Highlight = true
		updateEditorContent(g)
	}
	return nil
}

func updateEditorContent(g *gocui.Gui) {
	refreshEditorLayout(g)
	item := guiData_.editorData_.items[guiData_.editorData_.curr]
	v, err := g.View("Editor")
	if err != nil {
		return
	}

	v.Title = item.Title + " - Input"
	v.Clear()
	v.SetCursor(0, 0)

	fmt.Fprint(v, item.DefaultValue)
	g.SetCurrentView("Editor")
	g.SetViewOnTop("Editor")
}

func updateEditorLayout(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		updateEditorContent(g)
		return nil
	})
}

func handleSave(g *gocui.Gui, v *gocui.View) error {
	item := guiData_.editorData_.items[guiData_.editorData_.curr]
	item.DataFunction(strings.TrimSpace(v.Buffer()))
	guiData_.editorData_.curr++
	if guiData_.editorData_.curr == len(guiData_.editorData_.items) {
		backState(g)
		return nil
	}
	updateEditorLayout(g)
	return nil
}
