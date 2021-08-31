package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupApp(app *tview.Application, options *Options) *tview.Application {
	app.EnableMouse(true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			goToMainPage(options)
		} else if event.Key() == tcell.KeyCtrlK {
			setSearchBarText("", options)
			setSearchBarFocus(options)
		} else if event.Key() == tcell.KeyCtrlF {
			setSearchBarText("/", options)
			setSearchBarFocus(options)
		} else if event.Key() == tcell.KeyCtrlQ {
			os.Exit(0)
		} else if event.Key() == tcell.KeyCtrlX {
			listServers(options)
		} else if event.Key() == tcell.KeyCtrlO {
			listMacros(options)
		} else if event.Key() == tcell.KeyCtrlN {
			err := options.CommandsHistory.GoToNextCommand()
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorOrange)
			} else {
				refreshPage(options, false)
			}
		} else if event.Key() == tcell.KeyCtrlP {
			err := options.CommandsHistory.GoToPreviousCommand()
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorOrange)
			} else {
				refreshPage(options, false)
			}
		} else if event.Key() == tcell.KeyCtrlR {
			toggleAutoRefresh(options)
		} else if event.Key() == tcell.KeyCtrlH {
			showHome(options)
			setSearchBarText("", options)
		}
		return event
	})
	return app
}
