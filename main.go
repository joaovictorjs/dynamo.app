package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("dynamo.app")

	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetContent(widget.NewLabel("Hello World!"))
	win.ShowAndRun()
}
