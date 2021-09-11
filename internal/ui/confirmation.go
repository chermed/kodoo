package ui

import (
	"github.com/rivo/tview"
)

func setupConfirmation(options *Options) *tview.Grid {
	grid := options.Confirmation
	grid.Clear()
	grid.SetColumns(0, 10, 0).SetRows(0, 10, 0)
	modal := getConfirmationModal(options)
	modal.SetTextColor(options.Skin.ModalFgColor)
	modal.SetBackgroundColor(options.Skin.ModalBackgroundColor)
	modal.SetButtonTextColor(options.Skin.TitleColor)
	modal.AddButtons([]string{"Apply", "Cancel"})
	grid.AddItem(modal, 1, 1, 1, 1, 0, 0, true)
	return grid
}
func getConfirmationModal(options *Options) *tview.Modal {
	return options.ConfirmationModal
}
func showConfirmationModal(options *Options, title string, msg string, fn func() error) (err error) {
	modal := getConfirmationModal(options)
	modal.SetText(msg)
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		options.Pages.SwitchToPage("main")
		if buttonLabel == "Apply" {
			err = fn()
		}
	})
	options.Pages.ShowPage("confirmation")
	options.App.SetFocus(modal)
	return err
}
