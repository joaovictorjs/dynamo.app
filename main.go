package main

import "fyne.io/fyne/v2/app"

func main() {
	application := app.New()
	application.Settings().SetTheme(newDynamoTheme())
	mainWindow := application.NewWindow("dynamo.app")

	gui := newGui(mainWindow)

	mainWindow.SetContent(gui.makeGUI())
	mainWindow.SetPadded(false)
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
