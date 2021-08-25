package ui

import (
	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
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
	metadata.SetBorder(true).SetTitle(" Metadata ")
	metadata.SetTitleColor(options.Skin.TitleColor)
	metadata.SetColumns(26, 3, 0)
	model, ids, err := getTableModelIDs(options, true)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	cmd := odoo.NewCommandIDs(model, ids)
	results, err := data.ReadMetadata(cmd, options.Config, options.OdooCfg)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	if len(results) == 0 {
		showInfo("No metadata found for this record", options, tcell.ColorRed)
		return
	}
	result := results[0]
	addField(options, metadata, 0, 0, "Model", cmd.Model, false, false)
	addField(options, metadata, 1, 0, "ID", result.ID, false, false)
	addField(options, metadata, 2, 0, "XML-ID", result.XMLID, false, false)
	addField(options, metadata, 3, 0, "No Update", result.NoUpdate, false, false)
	addField(options, metadata, 4, 0, "Creation User", result.CreateUID, false, false)
	addField(options, metadata, 5, 0, "Creation Date", result.CreateDate, false, false)
	addField(options, metadata, 6, 0, "Latest Modification by", result.WriteUID, false, false)
	addField(options, metadata, 7, 0, "Latest Modification Date", result.WriteDate, false, false)
	options.Pages.ShowPage("metadata")
}
