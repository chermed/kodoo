package ui

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupSearchBar(grid *tview.Grid, options *Options) *tview.Grid {
	searchBar := grid.SetRows(1).SetColumns(0)
	inputField := options.InputField
	searchBar.SetBorder(true)
	searchBar.SetBorderColor(options.Skin.BorderColor)
	searchBar.SetBackgroundColor(options.Skin.BackgroundColor)
	inputField.SetBackgroundColor(options.Skin.BackgroundColor)
	inputField.SetFieldBackgroundColor(options.Skin.BackgroundColor)
	inputField.SetText("")
	inputField.SetFieldTextColor(options.Skin.FieldValueColor)
	inputField.SetLabel("QUERY> ")
	inputField.SetLabelColor(options.Skin.FieldLabelColor)
	searchBar.AddItem(inputField, 0, 0, 1, 1, 0, 0, true)
	searchBar.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		inputValue := strings.Trim(inputField.GetText(), " ")
		if strings.HasPrefix(inputValue, "#") {
			inputField.SetLabel("FIELDS> ")
		} else if strings.HasPrefix(inputValue, ">") {
			inputField.SetLabel("QUERY> ")
		} else if strings.HasPrefix(inputValue, ":") {
			inputField.SetLabel("COMMAND> ")
		} else if strings.HasPrefix(inputValue, "@") {
			inputField.SetLabel("MACRO> ")
		} else if strings.HasPrefix(inputValue, "/") {
			inputField.SetLabel("FILTER> ")
		} else if strings.HasPrefix(inputValue, "!") {
			inputField.SetLabel("CALL FUNC> ")
		} else {
			inputField.SetLabel("QUERY> ")
		}
		if event.Key() == tcell.KeyEnter {
			err := processInput(inputValue, options)
			if err != nil {
				showError(err, options)
			} else {
				options.QueryHistory.AddQuery(data.Query(inputField.GetText()))
				inputField.SetText("")
			}
		} else if event.Key() == tcell.KeyUp {
			if err := options.QueryHistory.GoToPreviousQuery(); err != nil {
				showInfo(err.Error(), options, tcell.ColorOrange)
			} else {
				query, err := options.QueryHistory.GetQuery()
				if err != nil {
					showInfo(err.Error(), options, tcell.ColorOrange)
				}
				inputField.SetText(string(query))
			}
		} else if event.Key() == tcell.KeyDown {
			if err := options.QueryHistory.GoToNextQuery(); err != nil {
				showInfo(err.Error(), options, tcell.ColorOrange)
			} else {
				query, err := options.QueryHistory.GetQuery()
				if err != nil {
					showInfo(err.Error(), options, tcell.ColorOrange)
				}
				inputField.SetText(string(query))
			}
		}
		return event
	})
	return searchBar
}

func processInput(value string, options *Options) error {
	CleanedValue := strings.Trim(value, " ")
	lowerCleanedValue := strings.ToLower(value)
	if CleanedValue == "" {
		err := errors.New("The input is empty")
		return showInfo(err.Error(), options, tcell.ColorRed)
	} else if lowerCleanedValue == "?" {
		showHome(options)
		return nil
	} else if lowerCleanedValue == ":q" {
		options.Shutdown <- "ok"
		os.Exit(0)
	} else if lowerCleanedValue == ":x" {
		return listServers(options)
	} else if lowerCleanedValue == ":o" {
		return listMacros(options)
	} else if strings.HasPrefix(CleanedValue, ":") {
		return fmt.Errorf("Command [%s] not found!", value)
	} else if strings.HasPrefix(CleanedValue, "!") {
		funcValue := strings.Split(CleanedValue[1:], " ")
		if len(funcValue) == 0 {
			return fmt.Errorf("Please specify a remote function to execute")
		}
		model, ids, err := getTableModelIDs(options)
		if err != nil {
			return showInfo(err.Error(), options, tcell.ColorRed)
		}
		funcName := funcValue[0]
		funcArgs := []interface{}{ids}
		if len(funcValue) > 1 {
			for _, strValue := range funcValue[1:] {
				funcArgs = append(funcArgs, strValue)
			}
		}
		title := fmt.Sprintf("Execute <%s> on %d item(s) ?", funcName, len(ids))
		question := fmt.Sprintf("Do you want to execute the function <%s> of the object <%s> on %d item(s) ?", funcName, model, len(ids))
		return showConfirmationModal(options, title, question, func() error {
			data, err := data.RunRemoteFunction(options.OdooCfg, options.Config, model, funcName, funcArgs...)
			if err != nil {
				return showInfo(err.Error(), options, tcell.ColorRed)
			} else {
				refreshPage(options, false)
				return showInfo(fmt.Sprintf("%+v", data.Result), options, tcell.ColorWhite)
			}
		})

	} else if strings.HasPrefix(CleanedValue, "/") {
		filterString := strings.Trim(CleanedValue[1:], " ")
		var cmd odoo.Command
		var err error
		if options.CommandsHistory.HasCommand() {
			cmd, err = options.CommandsHistory.GetCommandCopy()
			if err != nil {
				return err
			}
		} else {
			cmd = odoo.Command{}
		}
		err = odoo.StringToCommand(&cmd, filterString)
		if err != nil {
			return err
		}
		options.CommandsHistory.AddCommand(&cmd)
		return showSearchReadResult(&cmd, options)
	} else if strings.HasPrefix(lowerCleanedValue, "@") {
		macroName := CleanedValue[1:]
		if macroName == "" {
			showInfo("You can specify a name of the macro directly, e.g: @macroName", options, tcell.ColorOrangeRed)
			return listMacros(options)
		}
		macroCommand, err := getCommandFromMacro(macroName, options)
		if err != nil {
			return err
		}
		options.CommandsHistory.AddCommand(macroCommand)
		return showSearchReadResult(macroCommand, options)
	} else if strings.HasPrefix(CleanedValue, "#") {
		modelName := CleanedValue[1:]
		if modelName == "" {
			lastCommand, err := options.CommandsHistory.GetCommand()
			if err == nil {
				modelName = lastCommand.Model
			}
		}
		if modelName == "" {
			err := showInfo("Please provide the name of the object", options, tcell.ColorOrangeRed)
			return err
		} else {
			return listFields(options, modelName)
		}
	} else {
		if strings.HasPrefix(CleanedValue, ">") {
			CleanedValue = strings.Trim(CleanedValue, ">")
			CleanedValue = strings.Trim(CleanedValue, " ")
		}
		domain := make([][]interface{}, 0)
		cmd := odoo.NewCommand(
			options.OdooCfg,
			"",
			domain,
			[]string{},
			options.OdooCfg.DefaultLimit,
			"",
			odoo.OdooContext{},
		)
		if err := odoo.StringToCommand(cmd, CleanedValue); err != nil {
			return showError(err, options)
		}
		options.CommandsHistory.AddCommand(cmd)
		return showSearchReadResult(cmd, options)

	}
	return fmt.Errorf("Command [%s] not found!", value)
}

func setSearchBarFocus(options *Options) {
	options.App.SetFocus(options.InputField)
}
func setSearchBarText(txt string, options *Options) {
	options.InputField.SetText(txt)
}
