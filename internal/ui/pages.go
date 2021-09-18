package ui

import (
	"github.com/gdamore/tcell/v2"
)

func goToNextPage(options *Options, rotate bool) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToNextPage(rotate) {
		return showInfo("you are on the last page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}
func goToPreviousPage(options *Options, rotate bool) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToPreviousPage(rotate) {
		return showInfo("you are on the first page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}

func goToFirstPage(options *Options, rotate bool) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToFirstPage(rotate) {
		return showInfo("you are on the first page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}
func goToLastPage(options *Options, rotate bool) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToLastPage(rotate) {
		return showInfo("you are on the last page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}

func refreshPage(options *Options, show bool) error {
	if show {
		err := showInfo("refreshing ...", options, tcell.ColorGreen)
		if err != nil {
			return err
		}
	}
	if options.Config.MetaConfig.ZenMode {
		err := goToNextPage(options, true)
		if err != nil {
			return showInfo(err.Error(), options, tcell.ColorRed)
		}
	}
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	err = showSearchReadResult(command, options)
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	options.App.ForceDraw()
	return err
}

func goToMainPage(options *Options) {
	options.Pages.SwitchToPage("main")
	if mainContainerHasTable(options) {
		setTableFocus(options)
	}
}
