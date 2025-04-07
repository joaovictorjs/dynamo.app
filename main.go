package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(newDynamoTheme())
	win := app.NewWindow("dynamo.app")

	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetContent(makeGUI())
	win.ShowAndRun()
}
