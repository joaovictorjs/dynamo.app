package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"github.com/sqweek/dialog"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(newDynamoTheme())
	win := app.NewWindow("dynamo.app")

	actualDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		dialog.Message("%s", err.Error())
		return
	}

	gui := &gui{
		win:              win,
		currentDirectory: binding.BindString(&actualDir),
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
