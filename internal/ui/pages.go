package ui

import (
	"github.com/gdamore/tcell/v2"
)

func goToNextPage(options *Options) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToNextPage() {
		return showInfo("You are on the last page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}
func goToPreviousPage(options *Options) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToPreviousPage() {
		return showInfo("You are on the first page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}

func goToFirstPage(options *Options) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToFirstPage() {
		return showInfo("You are on the first page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}
func goToLastPage(options *Options) error {
	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	if !command.GoToLastPage() {
		return showInfo("You are on the last page", options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}

func refreshPage(options *Options, show bool) error {
	if show {
		err := showInfo("Refreshing ...", options, tcell.ColorGreen)
		if err != nil {
			return err
		}
	}

	command, err := options.CommandsHistory.GetCommand()
	if err != nil {
		return showInfo(err.Error(), options, tcell.ColorRed)
	}
	return showSearchReadResult(command, options)
}

func goToMainPage(options *Options) {
	options.Pages.SwitchToPage("main")
	if mainContainerHasTable(options) {
		setTableFocus(options)
	}
}
