package ui

import (
	"github.com/chermed/kodoo/internal/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getServerInfo(options *Options) *tview.Grid {
	infos := tview.NewGrid().SetRows(1, 1, 1, 1).SetColumns(9, 2, 0)
	currentConfig, err := data.GetCurrentServer(options.Config)
	if err != nil {
		addField(options, infos, 0, 0, "Error", err.Error(), true, true)
	} else {
		addField(options, infos, 0, 0, "Host", currentConfig.Host, false, true)
		addField(options, infos, 1, 0, "Database", currentConfig.Database, false, true)
		addField(options, infos, 2, 0, "User", currentConfig.User, false, true)
		if options.Config.MetaConfig.NoPassword {
			addField(options, infos, 3, 0, "Password", "*****", false, true)
		} else {
			addField(options, infos, 3, 0, "Password", currentConfig.Password, false, true)
		}
	}
	return infos
}

func listServers(options *Options) error {
	tableData := data.GetServers(*options.Config)
	table := getTableScreen(tableData, options)
	clearMainContainer(options)
	options.Table = table
	options.MainContainer.AddItem(table, 0, 0, 1, 1, 0, 0, true)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = checkTableSearchBarShortcuts(event, options)
		if event.Key() == tcell.KeyEnter {
			row, _ := table.GetSelection()
			options.Config.MetaConfig.DefaultServer = tableData.Lines[row-1]["name"].Value.(string)
			setupHeader(options.Header, options)
			clearMainContainer(options)
			setEmptyTextView(options)
			setSearchBarFocus(options)
		}
		return event
	})
	setTableFocus(options)
	return nil
}
