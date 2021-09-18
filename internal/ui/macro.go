package ui

import (
	"fmt"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
)

func listMacros(options *Options) error {
	tableData := data.GetMacros(*options.Config)
	table := getTableScreen(tableData, options)
	clearMainContainer(options)
	options.Table = table
	options.MainContainer.AddItem(table, 0, 0, 1, 1, 0, 0, true)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = checkTableSearchBarShortcuts(event, options)
		if event.Key() == tcell.KeyEnter {
			row, _ := table.GetSelection()
			macroName := tableData.Lines[row-1]["name"].Value.(string)
			macroCommand, err := getCommandFromMacro(macroName, options)
			if err != nil {
				showError(err, options)
			} else {
				if err := showSearchReadResult(macroCommand, options); err != nil {
					showError(err, options)
				} else {
					options.CommandsHistory.AddCommand(macroCommand)
				}
			}
		}
		return event
	})
	setTableFocus(options)
	return nil
}

func getCommandFromMacro(macroName string, options *Options) (*odoo.Command, error) {
	macro, found := options.Config.Macros[macroName]
	if !found {
		err := fmt.Errorf("the macro [%s] not found", macroName)
		return &odoo.Command{}, err
	}
	command := odoo.NewCommand(options.OdooCfg, macro.Model, macro.Domain, macro.Fields, macro.Limit, macro.Order, macro.Context)
	command.Description = macro.Description
	return command, nil
}
