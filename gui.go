package main

import (
	"errors"
	"image/color"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type gui struct {
	win              fyne.Window
	currentDirectory binding.String
}

func (g *gui) makeLeftPanel() fyne.CanvasObject {
	data := binding.BindStringList(dummyReportNames)

	list := widget.NewListWithData(
		data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		},
	)

	list.HideSeparators = true

	return container.NewPadded(list)
}

func (g *gui) makeTopBar() fyne.CanvasObject {
	var toolbarActionRef *widget.ToolbarAction
	currApp := fyne.CurrentApp()

	items := []*fyne.MenuItem{
		fyne.NewMenuItem("Abrir diretório", g.openDirectoryDialog),
		fyne.NewMenuItem("Abrir relatório", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Fechar", currApp.Quit),
	}

	popUpMenu := widget.NewPopUpMenu(
		fyne.NewMenu("MainMenu", items...),
		g.win.Canvas(),
	)

	toolbarAction := widget.NewToolbarAction(
		theme.Icon(theme.IconNameMenu),
		func() {
			if toolbarActionRef == nil {
				return
			}

			pos := currApp.Driver().AbsolutePositionForObject(toolbarActionRef.ToolbarObject())
			newPos := fyne.NewPos(pos.X, pos.Y+toolbarActionRef.ToolbarObject().MinSize().Height)

			popUpMenu.ShowAtPosition(newPos)
		},
	)

	toolbarActionRef = toolbarAction

	return container.NewPadded(widget.NewToolbar(toolbarAction))
}

func (g *gui) makeRightPanel() fyne.CanvasObject {
	return canvas.NewRectangle(color.RGBA{18, 18, 18, 255})
}

func (g *gui) makeBottomPanel() fyne.CanvasObject {
	return container.NewVBox()
}

func (g *gui) makeGUI() fyne.CanvasObject {
	dividers := [3]fyne.Widget{
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
	}

	objects := []fyne.CanvasObject{
		g.makeLeftPanel(),
		g.makeTopBar(),
		g.makeRightPanel(),
		g.makeBottomPanel(),
		dividers[0],
		dividers[1],
		dividers[2],
	}

	layout := newDynamoLayout(
		objects[0],
		objects[1],
		objects[2],
		objects[3],
		dividers,
	)

	return container.New(layout, objects...)
}

func (g *gui) openDirectoryDialog() {
	dir, err := dialog.Directory().Browse()

	if err != nil && !errors.Is(err, dialog.Cancelled) {
		dialog.Message("%s", err.Error()).Error()
		return
	}

	abs, err := filepath.Abs(dir)

	if err != nil {
		dialog.Message("%s", err.Error()).Error()
		return
	}

	g.currentDirectory.Set(abs)

}
