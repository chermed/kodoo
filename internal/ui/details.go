package ui

import (
	"fmt"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
	"github.com/kyokomi/emoji"
	"github.com/rivo/tview"
)

func setupDetails(options *Options) *tview.Grid {
	grid := options.Details
	grid.Clear()
	grid.SetBorder(true)
	grid.SetBorderColor(options.Skin.BorderColor)
	grid.SetBackgroundColor(options.Skin.BackgroundColor)
	grid.SetColumns(0).SetRows(0)
	grid.AddItem(getEmptyTextView(options), 0, 0, 1, 1, 0, 0, true)
	return grid
}

func showDetails(options *Options) {
	lastCommand, err := options.CommandsHistory.GetCommand()
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	model, ids, err := getTableModelIDs(options, true)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	detailsGrid := setupDetails(options)
	detailsGrid.Clear()
	allFieldsText := ""
	cmd := odoo.NewCommandIDs(model, ids)
	cmd.AllFields = lastCommand.AllFields
	cmd.UseAllFields()
	cmd.SetFieldsUpdated()
	data, err := data.Read(cmd, ids, options.Config, options.OdooCfg)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	if len(data) == 0 {
		showInfo("No data found in the server", options, tcell.ColorRed)
		return
	}
	name := data[0]["name"]
	title := fmt.Sprintf(" [red]%s [cyan](id=%d) [yellow]%s ", model, ids[0], name)
	detailsGrid.SetTitle(title)
	for _, fieldName := range cmd.Fields {
		fieldSpec, found := cmd.AllFields[fieldName]
		if !found {
			continue
		}
		fieldValue, found := data[0][fieldName]
		if !found {
			continue
		}
		value := fieldValue
		if fieldSpec.Type == "binary" {
			value = emoji.Sprint("📚")
		}
		allFieldsText += fmt.Sprintf("[red]%v[cyan] (%v)[red]: [white]%v\n", fieldName, fieldSpec.Description, value)
	}
	allFieldsTextView := tview.NewTextView()
	allFieldsTextView.SetText(allFieldsText)
	allFieldsTextView.SetDynamicColors(true)
	allFieldsTextView.SetBackgroundColor(options.Skin.BackgroundColor)
	detailsGrid.AddItem(allFieldsTextView, 0, 0, 1, 1, 0, 0, true)
	options.Pages.ShowPage("details")
}
