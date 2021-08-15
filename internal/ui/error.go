package ui

import (
	"github.com/rivo/tview"
)

func showError(err error, options *Options) error {
	if err != nil {
		clearMainContainer(options)
		errorArea := tview.NewTextView().SetText(err.Error())
		errorArea.SetBorder(true)
		errorArea.SetTextColor(options.Skin.FgErrorColor)
		errorArea.SetBorderColor(options.Skin.FgErrorColor)
		errorArea.SetBackgroundColor(options.Skin.BackgroundColor)
		errorArea.SetTextAlign(tview.AlignCenter)
		options.MainContainer.AddItem(errorArea, 0, 0, 1, 1, 0, 0, false)
		setSearchBarFocus(options)
	}
	return nil
}
