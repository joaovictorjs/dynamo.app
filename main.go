package main

import "fyne.io/fyne/v2/app"

func main() {
	application := app.New()
	mainWindow := application.NewWindow("dynamo.app")

	mainWindow.SetContent(makeGUI())
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
