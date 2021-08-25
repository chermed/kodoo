package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	logo = `	____            ____             
	|  | ______   __| _/____   ____  
	|  |/ /  _ \ / __ |/  _ \ /  _ \ 
	|    (  (_) ) /_/ (  (_) |  (_) )
	|__|_ \____/\____ |\____/ \____/ 
		 \/          \/              `

	note = `[yellow]Shortcuts are working in two modes : 
[yellow][red]1) Global[yellow]: see global shortcuts
[yellow][red]2) Table[yellow]: see table shortcuts
 
[yellow][pink] * [yellow]To switch from table mode to global mode: type one of the keys: [red]> ! : /[yellow] or [red]#[yellow]
 
[yellow][pink] * [yellow]To switch from global mode to table mode: type the key [red]Escape[yellow]`
)

func getLogo(options *Options) *tview.Grid {
	logoTextView := tview.NewTextView()
	logoTextView.SetText(logo)
	logoTextView.SetTextColor(options.Skin.TitleColor)
	logoTextView.SetBackgroundColor(options.Skin.BackgroundColor)
	logoGrid := tview.NewGrid().
		SetRows(6).
		SetColumns(0)
	logoGrid.AddItem(logoTextView, 0, 0, 1, 1, 0, 0, false)
	return logoGrid
}

func getInfos(options *Options) *tview.Grid {
	infoGrid := tview.NewGrid().
		SetRows(1, 1, 1, 1).
		SetColumns(8, 3, 0)
	infoGrid.SetBorder(true).SetTitle(" Informations ")
	infoGrid.SetBackgroundColor(options.Skin.BackgroundColor)
	infoGrid.SetBorderColor(options.Skin.BorderColor)
	infoGrid.SetTitleColor(options.Skin.TitleColor)
	addField(options, infoGrid, 0, 0, "URL", options.Config.URL, false, false)
	addField(options, infoGrid, 1, 0, "Version", options.Config.Version, false, false)
	addField(options, infoGrid, 2, 0, "", "", false, false)
	addField(options, infoGrid, 3, 0, "", "", false, false)
	return infoGrid
}
func getNote(options *Options) *tview.Grid {
	noteTextView := tview.NewTextView()
	noteTextView.SetText(note).SetDynamicColors(true)
	noteTextView.SetBackgroundColor(options.Skin.BackgroundColor)
	noteGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0)
	noteGrid.SetBorder(true).SetTitle(" Note ")
	noteGrid.SetBackgroundColor(options.Skin.BackgroundColor)
	noteGrid.SetBorderColor(options.Skin.BorderColor)
	noteGrid.SetTitleColor(options.Skin.TitleColor)
	noteGrid.AddItem(noteTextView, 0, 0, 1, 1, 0, 0, false)
	return noteGrid
}
func getQuickActions(options *Options) *tview.Grid {
	quickActionsGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(13, 3, 0)
	quickActionsGrid.SetBorder(true).SetTitle(" Quick Actions ")
	quickActionsGrid.SetBackgroundColor(options.Skin.BackgroundColor)
	quickActionsGrid.SetBorderColor(options.Skin.BorderColor)
	quickActionsGrid.SetTitleColor(options.Skin.TitleColor)
	addField(options, quickActionsGrid, 0, 0, "?", "Show help", false, false)
	addField(options, quickActionsGrid, 1, 0, ":q", "Exit the application", false, false)
	addField(options, quickActionsGrid, 2, 0, ":x", "List the servers", false, false)
	addField(options, quickActionsGrid, 3, 0, ":o", "List the macros", false, false)
	addField(options, quickActionsGrid, 4, 0, "@MacroName", "Execute the macro named <MacroName>", false, false)
	addField(options, quickActionsGrid, 5, 0, "#Object", "Show the fields of the object <Object>", false, false)
	addField(options, quickActionsGrid, 6, 0, ">Query", "Run a query (e.g: >sale.order id +name state state=draft partner_id.name~azure)", false, false)
	addField(options, quickActionsGrid, 7, 0, "", "", false, false)
	return quickActionsGrid
}
func getGlobalShortcuts(options *Options) *tview.Grid {
	globalShortcutsGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(13, 3, 0)
	globalShortcutsGrid.SetBorder(true).SetTitle(" Global shortcuts ")
	globalShortcutsGrid.SetBackgroundColor(options.Skin.BackgroundColor)
	globalShortcutsGrid.SetBorderColor(options.Skin.BorderColor)
	globalShortcutsGrid.SetTitleColor(options.Skin.TitleColor)
	addField(options, globalShortcutsGrid, 0, 0, "Ctrl-K", "Clear the input", false, false)
	addField(options, globalShortcutsGrid, 1, 0, "Ctrl-N", "Go to the next command", false, false)
	addField(options, globalShortcutsGrid, 2, 0, "Ctrl-P", "Go to the previous command", false, false)
	addField(options, globalShortcutsGrid, 3, 0, "Ctrl-F", "Filter records", false, false)
	addField(options, globalShortcutsGrid, 4, 0, "Ctrl-O", "List macros", false, false)
	addField(options, globalShortcutsGrid, 5, 0, "Ctrl-R", "Enable or disable auto-refresh", false, false)
	addField(options, globalShortcutsGrid, 6, 0, "Ctrl-X", "List servers", false, false)
	addField(options, globalShortcutsGrid, 7, 0, "Ctrl-H", "Show and hide the help", false, false)
	addField(options, globalShortcutsGrid, 8, 0, "Ctrl-Q", "Exit the application", false, false)
	addField(options, globalShortcutsGrid, 9, 0, "Key UP/DOWN", "Navigate through command history", false, false)
	return globalShortcutsGrid
}
func getTableShortcuts(options *Options) *tview.Grid {
	tableShortcutsGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(13, 3, 0)
	tableShortcutsGrid.SetBorder(true).SetTitle(" Table shortcuts ")
	tableShortcutsGrid.SetBackgroundColor(options.Skin.BackgroundColor)
	tableShortcutsGrid.SetBorderColor(options.Skin.BorderColor)
	tableShortcutsGrid.SetTitleColor(options.Skin.TitleColor)
	addField(options, tableShortcutsGrid, 0, 0, "n", "go to the next page", false, false)
	addField(options, tableShortcutsGrid, 1, 0, "p", "go to the previous page", false, false)
	addField(options, tableShortcutsGrid, 2, 0, "l", "go to the last page", false, false)
	addField(options, tableShortcutsGrid, 3, 0, "f", "go to the first page", false, false)
	addField(options, tableShortcutsGrid, 4, 0, "r", "refresh the data", false, false)
	addField(options, tableShortcutsGrid, 5, 0, "Key <ENTER>", "open related records", false, false)
	addField(options, tableShortcutsGrid, 6, 0, "!FuncName", "execute the remote function <FuncName>", false, false)
	addField(options, tableShortcutsGrid, 7, 0, "/filter", "filter the records", false, false)
	addField(options, tableShortcutsGrid, 8, 0, "", "", false, false)
	addField(options, tableShortcutsGrid, 9, 0, "", "", false, false)
	return tableShortcutsGrid
}
func getHelpFooter(options *Options) *tview.Grid {
	helpFooterGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0)
	helpFooterTextView := tview.NewTextView()
	const code string = `
	[green]Type the shortcut [red]Ctrl-H[green] or [red]q[green] to quit the help mode[green]
	[green]You can also type one of the keys: [red]> @ :[green] or [red]#[green]`
	helpFooterTextView.SetText(code)
	helpFooterTextView.SetTextAlign(tview.AlignCenter)
	helpFooterTextView.SetDynamicColors(true)
	helpFooterTextView.SetBackgroundColor(options.Skin.BackgroundColor)
	helpFooterGrid.AddItem(helpFooterTextView, 0, 0, 1, 1, 0, 0, false)
	return helpFooterGrid
}
func getHomeGrid(options *Options) *tview.Grid {
	homeGrid := tview.NewGrid().
		SetRows(6, 11, 12, 0).
		SetColumns(0, 0)
	homeGrid.AddItem(getLogo(options), 0, 0, 1, 1, 0, 0, false)
	homeGrid.AddItem(getInfos(options), 0, 1, 1, 1, 0, 0, false)
	homeGrid.AddItem(getNote(options), 1, 0, 1, 1, 0, 0, false)
	homeGrid.AddItem(getQuickActions(options), 1, 1, 1, 1, 0, 0, false)
	homeGrid.AddItem(getGlobalShortcuts(options), 2, 0, 1, 1, 0, 0, false)
	homeGrid.AddItem(getTableShortcuts(options), 2, 1, 1, 1, 0, 0, false)
	homeGrid.AddItem(getHelpFooter(options), 3, 0, 1, 2, 0, 0, false)
	homeGrid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, r := range []rune{'>', '@', ':', '#'} {
			if event.Rune() == r {
				setSearchBarText(string(r), options)
				options.Pages.SwitchToPage("main")
				setSearchBarFocus(options)
			}
		}
		if event.Rune() == 'q' {
			options.Pages.SwitchToPage("main")
		}
		return event
	})
	return homeGrid
}
func showHome(options *Options) {
	pageName, _ := options.Pages.GetFrontPage()
	if pageName == "home" {
		options.Pages.SwitchToPage("main")
	} else {
		options.Pages.SwitchToPage("home")
	}
}
