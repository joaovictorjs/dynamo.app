package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeTopBar() fyne.CanvasObject {
	var toolbarActionRef *widget.ToolbarAction

	items := []*fyne.MenuItem{
		fyne.NewMenuItem("Abrir diretório", func() {}),
		fyne.NewMenuItem("Abrir relatório", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Fecha", func() {}),
	}

	curApp := fyne.CurrentApp()

	popUpMenu := widget.NewPopUpMenu(
		fyne.NewMenu("MainMenu", items...),
		curApp.Driver().AllWindows()[0].Canvas(),
	)

	toolbarAction := widget.NewToolbarAction(
		theme.Icon(theme.IconNameMenu),
		func() {
			if toolbarActionRef == nil {
				return
			}

			pos := curApp.Driver().AbsolutePositionForObject(toolbarActionRef.ToolbarObject())
			newPos := fyne.NewPos(0, pos.Y+toolbarActionRef.ToolbarObject().MinSize().Height)

			popUpMenu.ShowAtPosition(newPos)
		},
	)

	toolbarActionRef = toolbarAction

	return container.NewPadded(widget.NewToolbar(toolbarAction))
}

func makeGUI() fyne.CanvasObject {
	dividers := [3]fyne.Widget{
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
	}

	objects := []fyne.CanvasObject{
		makeTopBar(),
		dividers[0],
		dividers[1],
		dividers[2],
	}

	layout := newDynamoLayout(
		objects[0],
		dividers,
	)

	return container.New(layout, objects...)
}
