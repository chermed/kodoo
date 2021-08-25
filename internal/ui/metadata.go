package ui

import (
	"github.com/chermed/kodoo/internal/data"
	"github.com/rivo/tview"
)

func setupMetadataGrid(options *Options) *tview.Grid {
	grid := options.MetadataGrid
	grid.Clear()
	grid.SetBackgroundColor(options.Skin.BackgroundColor)
	grid.SetColumns(0, 100, 0).SetRows(0, 10, 0)
	grid.AddItem(getMetadata(options), 1, 1, 1, 1, 0, 0, true)
	return grid
}
func getMetadata(options *Options) *tview.Grid {
	return options.Metadata
}
func showMetadata(options *Options) {
	metadata := getMetadata(options)
	metadata.Clear()
	metadata.SetBorder(true).SetTitle("Metadata")
	metadata.SetColumns(26, 3, 0)
	// model, ids, err := getTableModelID(Options)
	data.ReadMetadata(options.Config, options.OdooCfg, options.CommandsHistory)
	addField(options, metadata, 0, 0, "Model", "ffff", false)
	addField(options, metadata, 1, 0, "ID", "3333", false)
	addField(options, metadata, 2, 0, "XML-ID", "3333", false)
	addField(options, metadata, 3, 0, "No Update", "3333", false)
	addField(options, metadata, 4, 0, "Creation User", "3333", false)
	addField(options, metadata, 5, 0, "Creation Date", "3333", false)
	addField(options, metadata, 6, 0, "Latest Modification by", "3333", false)
	addField(options, metadata, 7, 0, "Latest Modification Date", "3333", false)
	options.Pages.ShowPage("metadata")
}
