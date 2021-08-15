package ui

import (
	"github.com/chermed/kodoo/internal/data"
	"github.com/gdamore/tcell/v2"
)

func listFields(options *Options, model string) error {
	tableData, err := data.GetFields(options.Config, options.OdooCfg, model)
	if err != nil {
		return err
	}
	table := getTableScreen(tableData, options)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = checkTableSearchBarShortcuts(event, options)
		return event
	})
	clearMainContainer(options)
	options.Table = table
	options.MainContainer.AddItem(table, 0, 0, 1, 1, 0, 0, true)
	setTableFocus(options)
	return nil
}
