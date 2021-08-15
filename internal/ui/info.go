package ui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func showInfo(msg string, options *Options, color tcell.Color) error {
	emptyMsg := " "
	if msg != "" {
		clearFooter(options)
		infoArea := tview.NewTextView().SetText(msg)
		infoArea.SetTextColor(color)
		infoArea.SetBackgroundColor(options.Skin.BackgroundColor)
		infoArea.SetTextAlign(tview.AlignCenter)
		options.Footer.AddItem(infoArea, 0, 0, 1, 1, 0, 0, false)
	}
	if msg != emptyMsg {
		timer := time.NewTimer(4 * time.Second)
		go func() {
			<-timer.C
			showInfo(emptyMsg, options, tcell.ColorRed)
			options.App.ForceDraw()
		}()
	}
	return nil
}
