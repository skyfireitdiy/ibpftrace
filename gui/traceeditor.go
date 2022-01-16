package gui

import (
	"github.com/jroimartin/gocui"
)

func syncTracerEditorInfo(g *gocui.Gui) {
	guiData_.editorData_.items = guiData_.tracerEditorData_.items
}

func refreshTracerEditorLayout(g *gocui.Gui) error {
	syncTracerEditorInfo(g)
	refreshEditorLayout(g)
	return nil
}

func updateTracerEditorLayout(g *gocui.Gui) {
	syncTracerEditorInfo(g)
	updateEditorLayout(g)
}

func saveTracerEditorViewData() interface{} {
	return guiData_.tracerEditorData_
}

func recoverTracerEditorViewData(info stateInfo) {
	guiData_.tracerEditorData_ = info.data.(tracerEditorViewData)
}
