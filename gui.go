package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func makeGUI() fyne.CanvasObject {
	return container.New(newResponsiveLayout(nil, nil, nil))
}
