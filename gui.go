package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeLeftPanel() fyne.CanvasObject {
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

func makeTopBar() fyne.CanvasObject {
	var toolbarActionRef *widget.ToolbarAction
	currApp := fyne.CurrentApp()

	items := []*fyne.MenuItem{
		fyne.NewMenuItem("Abrir diretório", func() {}),
		fyne.NewMenuItem("Abrir relatório", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Fechar", currApp.Quit),
	}

	popUpMenu := widget.NewPopUpMenu(
		fyne.NewMenu("MainMenu", items...),
		currApp.Driver().AllWindows()[0].Canvas(),
	)

	toolbarAction := widget.NewToolbarAction(
		theme.Icon(theme.IconNameMenu),
		func() {
			if toolbarActionRef == nil {
				return
			}

			pos := currApp.Driver().AbsolutePositionForObject(toolbarActionRef.ToolbarObject())
			newPos := fyne.NewPos(0, pos.Y+toolbarActionRef.ToolbarObject().MinSize().Height)

			popUpMenu.ShowAtPosition(newPos)
		},
	)

	toolbarActionRef = toolbarAction

	return container.NewPadded(widget.NewToolbar(toolbarAction))
}

func makeRightPanel() fyne.CanvasObject {
	return canvas.NewRectangle(color.RGBA{18, 18, 18, 255})
}

func makeGUI() fyne.CanvasObject {
	dividers := [3]fyne.Widget{
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
	}

	objects := []fyne.CanvasObject{
		makeLeftPanel(),
		makeTopBar(),
		makeRightPanel(),
		dividers[0],
		dividers[1],
		dividers[2],
	}

	layout := newDynamoLayout(
		objects[0],
		objects[1],
		objects[2],
		dividers,
	)

	return container.New(layout, objects...)
}
