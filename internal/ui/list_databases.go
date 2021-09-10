package ui

import (
	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/kotils"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	lblDatabases = "Databases"
	lblLogin     = "Username / Email"
	lblPassword  = "Password"
)

func setupListDatabasesGrid(options *Options) *tview.Grid {
	form := options.ListDatabasesForm
	form.AddDropDown(lblDatabases, []string{}, 0, nil)
	form.AddInputField(lblLogin, "", 30, nil, nil)
	form.AddPasswordField(lblPassword, "", 30, '*', nil)
	form.AddButton("Validate", func() {
		server, err := data.GetCurrentServer(options.Config)
		if err != nil {
			showInfo(err.Error(), options, tcell.ColorRed)
			return
		}
		dbsFormItem := form.GetFormItemByLabel(lblDatabases).(*tview.DropDown)
		loginFormItem := form.GetFormItemByLabel(lblLogin).(*tview.InputField)
		passwordFormItem := form.GetFormItemByLabel(lblPassword).(*tview.InputField)
		host := server.Host
		_, database := dbsFormItem.GetCurrentOption()
		user := loginFormItem.GetText()
		password := passwordFormItem.GetText()
		newServer := odoo.NewServer(host, database, user, password)
		if user == "" || password == "" || database == "" {
			showInfo("Some fields are empty", options, tcell.ColorRed)
			return
		}
		newServerName := kotils.RandSeq(10)
		options.Config.Servers[newServerName] = newServer
		data.SetCurrentServer(options.Config, newServerName)
		setupHeader(options.Header, options)
		clearMainContainer(options)
		setEmptyTextView(options)
		setSearchBarFocus(options)
		goToMainPage(options)
	})
	form.SetButtonsAlign(tview.AlignCenter)
	form.SetBorder(true)
	form.SetTitle(" Change the database / User ")
	form.SetBackgroundColor(options.Skin.BackgroundColor)
	grid := options.ListDatabasesGrid
	grid.Clear()
	grid.SetBackgroundColor(options.Skin.BackgroundColor)
	grid.SetColumns(0, 60, 0).SetRows(0, 12, 0)
	grid.AddItem(form, 1, 1, 1, 1, 0, 0, true)
	return grid
}

func showListDatabases(options *Options) {
	listDatabasesForm := options.ListDatabasesForm
	dbsFormItem := listDatabasesForm.GetFormItemByLabel(lblDatabases).(*tview.DropDown)
	loginFormItem := listDatabasesForm.GetFormItemByLabel(lblLogin).(*tview.InputField)
	passwordFormItem := listDatabasesForm.GetFormItemByLabel(lblPassword).(*tview.InputField)
	dbs, err := data.ListDatabases(options.Config, options.OdooCfg)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	server, err := data.GetCurrentServer(options.Config)
	if err != nil {
		showInfo(err.Error(), options, tcell.ColorRed)
		return
	}
	dbsFormItem.SetOptions(dbs, nil)
	dbsFormItem.SetCurrentOption(0)
	loginFormItem.SetText(server.User)
	passwordFormItem.SetText(server.Password)
	options.Pages.ShowPage("listDatabases")
}
