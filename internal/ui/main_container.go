package ui

import (
	"github.com/rivo/tview"
)

func setupMainContainer(grid *tview.Grid, options *Options) *tview.Grid {
	grid.Clear()
	mainContainer := grid.SetRows(0).SetColumns(0)
	return mainContainer
}

func clearMainContainer(options *Options) {
	options.MainContainer.Clear()
	removeTable(options)
}
func setEmptyTextView(options *Options) {
	textView := getEmptyTextView(options)
	options.MainContainer.AddItem(textView, 0, 0, 1, 1, 0, 0, false)
}
