package ui

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/rivo/tview"
)

type Options struct {
	App               *tview.Application
	Header            *tview.Grid
	Footer            *tview.Grid
	Details           *tview.Grid
	MetadataGrid      *tview.Grid
	Metadata          *tview.Grid
	ShortcutsGrid     *tview.Grid
	Shortcuts         *tview.List
	ListDatabasesGrid *tview.Grid
	ListDatabasesForm *tview.Form
	ListDatabases     *tview.List
	MainContainer     *tview.Grid
	Layout            *tview.Grid
	Config            *config.Config
	SearchBar         *tview.Grid
	InputField        *tview.InputField
	CommandsHistory   *data.CommandsHistory
	QueryHistory      *data.QueryHistory
	OdooCfg           *odoo.OdooConfig
	Table             *tview.Table
	Pages             *tview.Pages
	ConfirmationModal *tview.Modal
	Confirmation      *tview.Grid
	Skin              config.Skin
	Shutdown          chan string
}

func AppRun(cfg config.Config) {
	app := tview.NewApplication()
	pages := tview.NewPages()
	header := tview.NewGrid()
	footer := tview.NewGrid()
	mainContainer := tview.NewGrid()
	searchBar := tview.NewGrid()
	inputField := tview.NewInputField()
	layout := tview.NewGrid()
	details := tview.NewGrid()
	metadata := tview.NewGrid()
	metadataGrid := tview.NewGrid()
	shortcuts := tview.NewList()
	shortcutsGrid := tview.NewGrid()
	listDatabases := tview.NewList()
	listDatabasesGrid := tview.NewGrid()
	listDatabasesForm := tview.NewForm()
	confirmationModal := tview.NewModal()
	confirmation := tview.NewGrid()
	if cfg.MetaConfig.DefaultLimit == 0 {
		cfg.MetaConfig.DefaultLimit = 50
	}
	options := &Options{
		App:               app,
		Header:            header,
		Footer:            footer,
		Details:           details,
		MetadataGrid:      metadataGrid,
		Metadata:          metadata,
		ShortcutsGrid:     shortcutsGrid,
		Shortcuts:         shortcuts,
		ListDatabasesGrid: listDatabasesGrid,
		ListDatabasesForm: listDatabasesForm,
		ListDatabases:     listDatabases,
		MainContainer:     mainContainer,
		Layout:            layout,
		Config:            &cfg,
		SearchBar:         searchBar,
		InputField:        inputField,
		ConfirmationModal: confirmationModal,
		Confirmation:      confirmation,
		CommandsHistory:   &data.CommandsHistory{},
		QueryHistory:      &data.QueryHistory{},
		Pages:             pages,
		OdooCfg: &odoo.OdooConfig{
			DefaultLimit: cfg.MetaConfig.DefaultLimit,
			Log:          cfg.Log,
			Timeout:      cfg.MetaConfig.Timeout,
		},
		Skin:     config.GetSkin(),
		Shutdown: make(chan string),
	}
	setupApp(app, options)
	resetLayout(options)
	home := getHomeGrid(options)
	pages.AddPage("main", layout, true, true)
	pages.AddPage("home", home, true, false)
	pages.AddPage("confirmation", confirmation, true, false)
	pages.AddPage("details", details, true, false)
	pages.AddPage("metadata", metadataGrid, true, false)
	pages.AddPage("shortcuts", shortcutsGrid, true, false)
	pages.AddPage("listDatabases", listDatabasesGrid, true, false)
	if cfg.MetaConfig.DefaultMacro != "" {
		macroCommand, err := getCommandFromMacro(cfg.MetaConfig.DefaultMacro, options)
		if err != nil {
			showError(err, options)
		} else {
			options.CommandsHistory.AddCommand(macroCommand)
			err = showSearchReadResult(macroCommand, options)
			if err != nil {
				showError(err, options)
			} else {
				setTableFocus(options)
			}
		}
	} else {
		showHome(options)
	}
	startRefreshTicker(options)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func resetLayout(options *Options) {
	setupLayout(options, options.Config.MetaConfig.NoHeader, false, false)
	clearMainContainer(options)
}
func setupLayout(options *Options, noHeader bool, noSearchBar bool, noFooter bool) {
	options.Layout.Clear()
	options.Layout.SetColumns(0)
	setupConfirmation(options)
	setupMetadataGrid(options)
	setupDetails(options)
	setupShortcutsGrid(options)
	setupListDatabasesGrid(options)
	if noHeader && noSearchBar && noFooter {
		options.Layout.SetRows(0)
	} else if noHeader && noSearchBar {
		options.Layout.SetRows(0, 1)
	} else if noHeader {
		options.Layout.SetRows(3, 0, 1)
	} else {
		options.Layout.SetRows(4, 3, 0, 1)
	}
	var rowIndex int
	if !noHeader {
		setupHeader(options.Header, options)
		options.Layout.AddItem(options.Header, rowIndex, 0, 1, 1, 0, 0, false)
		rowIndex++
	}
	if !noSearchBar {
		setupSearchBar(options.SearchBar, options)
		options.Layout.AddItem(options.SearchBar, rowIndex, 0, 1, 1, 0, 0, true)
		rowIndex++
	}
	setupMainContainer(options.MainContainer, options)
	options.Layout.AddItem(options.MainContainer, rowIndex, 0, 1, 1, 0, 0, false)
	if !noFooter {
		rowIndex++
		setupFooter(options.Footer, options)
		options.Layout.AddItem(options.Footer, rowIndex, 0, 1, 1, 0, 0, false)
	}
}
