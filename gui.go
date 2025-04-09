package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	currentWindow fyne.Window
}

func newGui(win fyne.Window) *gui {
	return &gui{
		currentWindow: win,
	}
}

func (g *gui) makeTopContent() fyne.CanvasObject {
	popupMenu := widget.NewPopUpMenu(
		fyne.NewMenu(
			"",
			fyne.NewMenuItem("Abrir Diretório", func() {}),
			fyne.NewMenuItem("Abrir Relatório", func() {}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Fechar", func() {}),
		),
		g.currentWindow.Canvas(),
	)

	var menuActionRef *widget.ToolbarAction

	menuAction := widget.NewToolbarAction(
		theme.Icon(theme.IconNameMenu),
		func() {
			if menuActionRef == nil {
				return
			}

			pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(menuActionRef.ToolbarObject())
			popupMenu.ShowAtPosition(fyne.NewPos(pos.X, pos.Y+menuActionRef.ToolbarObject().MinSize().Height))
		},
	)

	menuActionRef = menuAction

	return container.NewPadded(
		container.NewHBox(
			widget.NewToolbar(menuAction),
			layout.NewSpacer(),
			widget.NewToolbar(
				widget.NewToolbarAction(theme.Icon(theme.IconNameMediaPlay), func() {}),
				widget.NewToolbarAction(theme.Icon(theme.IconNameMediaStop), func() {}),
				widget.NewToolbarAction(theme.Icon(theme.IconNameDownload), func() {}),
			),
		),
	)
}

func (g *gui) makeLeftContent() fyne.CanvasObject {
	return container.NewCenter()
}

func (g *gui) makeRightContent() fyne.CanvasObject {
	return container.NewCenter()
}

func (g *gui) makeCenterContent() fyne.CanvasObject {
	split := container.NewHSplit(
		g.makeLeftContent(),
		g.makeRightContent(),
	)

	split.Offset = 0.25

	return split
}

func (g *gui) makeBottomContent() fyne.CanvasObject {
	return container.NewCenter()
}

func (g *gui) makeGUI() fyne.CanvasObject {
	split := container.NewVSplit(
		g.makeCenterContent(),
		g.makeBottomContent(),
	)

	split.Offset = 0.65

	objects := []fyne.CanvasObject{
		g.makeTopContent(),
		widget.NewSeparator(),
		split,
	}

	return container.New(
		newDynamoLayout(
			objects[0],
			objects[1],
			objects[2],
		),
		objects...,
	)
}
