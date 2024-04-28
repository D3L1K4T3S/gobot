package gui

import (
	g "gobot/gui/windows"
	"gobot/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp() {
	state := state.NewAppState()

	app := app.New()
	icon, err := fyne.LoadResourceFromPath("../../assets/icon.png")
	if err != nil {
		panic(err)
	}

	app.SetIcon(icon)

	mainWindow := app.NewWindow("Main")
	startWindow := app.NewWindow("Start")

	mainWindow.SetMaster()
	mainWindow.SetFixedSize(true)
	mainWindow.SetCloseIntercept(func() {
		app.Quit()
	})

	startWindow.SetFixedSize(true)
	startWindow.SetCloseIntercept(func() {
		app.Quit()
	})

	g.StartRender(&app, state, startWindow, mainWindow)

	startWindow.Resize(fyne.NewSize(320, 160))
	startWindow.Show()
	mainWindow.Hide()
	app.Run()
}
