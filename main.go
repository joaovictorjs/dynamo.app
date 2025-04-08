package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"github.com/sqweek/dialog"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(newDynamoTheme())
	win := app.NewWindow("dynamo.app")

	gui := &gui{
		win:              win,
		currentDirectory: binding.NewString(),
	}

	gui.currentDirectory.AddListener(
		binding.NewDataListener(func() {
			dir, err := gui.currentDirectory.Get()

			if err != nil {
				dialog.Message("%s", err.Error())
				return
			}

			if dir == "" {
				return
			}

			win.SetTitle("dynamo.app: " + dir)
		}),
	)

	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetContent(gui.makeGUI())
	win.ShowAndRun()
}
