package ui

import (
	"github.com/rivo/tview"
)

func setupHeader(grid *tview.Grid, options *Options) *tview.Grid {
	grid.Clear()
	var header *tview.Grid
	if options.Config.MetaConfig.NoHeader {
		options.Layout.SetRows(0)
		header = grid.SetRows(0).SetColumns(0)
	} else {
		infos := getServerInfo(options)
		header = grid.SetRows(4).SetColumns(1, 0)
		header.SetBackgroundColor(options.Skin.BackgroundColor)
		header.AddItem(infos, 0, 1, 1, 1, 0, 0, false)
	}
	return header
}
