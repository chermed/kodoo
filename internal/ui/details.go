package ui

import "github.com/rivo/tview"

func setupDetails(options *Options) *tview.Grid {
	grid := options.Details
	grid.Clear()
	grid.SetColumns(0, 10, 0).SetRows(0, 10, 0)
	grid.AddItem(getEmptyTextView(options), 0, 0, 1, 1, 0, 0, true)
	return grid
}

func showDetails(options *Options) {
	options.Pages.ShowPage("details")
}
