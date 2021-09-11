package ui

import (
	"fmt"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupShortcutsGrid(options *Options) *tview.Grid {
	grid := options.ShortcutsGrid
	grid.Clear()
	grid.SetBackgroundColor(options.Skin.BackgroundColor)
	grid.SetColumns(0, 80, 0).SetRows(0, 30, 0)
	grid.AddItem(getShortcuts(options), 1, 1, 1, 1, 0, 0, true)
	return grid
}
func getShortcuts(options *Options) *tview.List {
	shortcuts := options.Shortcuts
	shortcuts.SetBackgroundColor(options.Skin.BackgroundColor)
	shortcuts.SetMainTextColor(options.Skin.SecondaryFgColor)
	shortcuts.SetSecondaryTextColor(options.Skin.OptionalFgColor)
	shortcuts.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			_, ids, err := getTableModelIDs(options, true)
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			_, model := shortcuts.GetItemText(shortcuts.GetCurrentItem())
			lastCommand, err := options.CommandsHistory.GetCommand()
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			rcmds, err := odoo.GetRelatedCommands(options.OdooCfg, lastCommand)
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			rcmd, err := odoo.GetRelatedCommand(model, rcmds)
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			rcmd.SetIDs(ids)
			cmd := rcmd.GetCommand(options.OdooCfg)
			if err = showSearchReadResult(cmd, options); err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			} else {
				options.CommandsHistory.AddCommand(cmd)
			}
			goToMainPage(options)
		}
		return event
	})
	return shortcuts
}
func showShortcuts(options *Options) {
	shortcuts := getShortcuts(options)
	shortcuts.Clear()
	shortcuts.SetBorder(true).SetTitle(" Shortcuts ")
	shortcuts.SetTitleColor(options.Skin.TitleColor)
	rcmds, err := data.GetRelatedCommands(options.CommandsHistory, options.OdooCfg)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	for _, rcmd := range rcmds {
		description := fmt.Sprintf("> %s (field: %s, type: %s)", rcmd.Description, rcmd.OriginField, rcmd.Type)
		shortcuts.AddItem(description, rcmd.Model, 0, nil)
	}
	options.Pages.ShowPage("shortcuts")
}
