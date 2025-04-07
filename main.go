package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(newDynamoTheme())
	win := app.NewWindow("dynamo.app")

	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetContent(container.NewVBox(widget.NewLabel("Hello World!"), widget.NewSeparator()))
	win.ShowAndRun()
}
