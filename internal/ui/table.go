package ui

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/kotils"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/kyokomi/emoji"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type OdooCellReference struct {
	FieldName string
	Model     string
}

func getX2ManyStrValue(options *Options, fieldName string, x2ManyData odoo.X2ManyResult, ids []interface{}) string {
	names := []string{}
	for _, id := range ids {
		names = append(names, x2ManyData[fieldName][int(id.(float64))])
	}
	return strings.Join(names, ", ")
}
func getTableScreen(tableData data.Data, options *Options) *tview.Table {
	log := options.Config.Log
	table := tview.NewTable().
		SetFixed(1, 1).SetSelectable(true, false)
	table.SetBackgroundColor(options.Skin.BackgroundColor)
	table.SetBorderColor(options.Skin.BorderColor)
	table.SetTitleColor(options.Skin.TitleColor)
	selectionTableCell := tview.NewTableCell("SEL").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter).
		SetSelectable(false).
		SetExpansion(0).
		SetBackgroundColor(options.Skin.BackgroundColor).
		SetTextColor(options.Skin.TableHeaderFgColor).
		SetMaxWidth(5)
	table.SetCell(0, 0, selectionTableCell)
	headerValueMap := make(map[int]string)
	for column, headerValue := range tableData.Header {
		headerValueMap[column] = headerValue
		tableCell := tview.NewTableCell(strings.ToUpper(headerValue)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignLeft).
			SetSelectable(false).
			SetExpansion(1).
			SetBackgroundColor(options.Skin.BackgroundColor).
			SetTextColor(options.Skin.TableHeaderFgColor)
		table.SetCell(0, column+1, tableCell)
	}
	idx := 0
	for row, lines := range tableData.Lines {
		fgColor := options.Skin.TableBodyFgColor
		selectionBodyTableCell := tview.NewTableCell("").
			SetAlign(tview.AlignCenter).
			SetSelectable(true).
			SetExpansion(0).
			SetBackgroundColor(options.Skin.BackgroundColor).
			SetTextColor(fgColor)
		if kotils.IntInSlice(idx, tableData.Selection) {
			selectionBodyTableCell.SetText("*")
			table.Select(row+1, 0)
		}
		table.SetCell(row+1, 0, selectionBodyTableCell)
		for column := range tableData.Header {
			tableCell := tview.NewTableCell("").
				SetAlign(tview.AlignLeft).
				SetSelectable(true).
				SetExpansion(1).
				SetBackgroundColor(options.Skin.BackgroundColor).
				SetTextColor(fgColor)
			fieldName := headerValueMap[column]
			odooCelleReference := &OdooCellReference{
				FieldName: fieldName,
				Model:     tableData.Model,
			}
			tableCell = tableCell.SetReference(odooCelleReference)
			item := lines[fieldName]
			strValue := ""
			switch item.Type {
			case "char", "string", "selection", "text":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				default:
					strValue = item.Value.(string)
					if fieldName == "state" {
						colorToApply, colorFound := options.Skin.StatesColor[strValue]
						if colorFound {
							fgColor = colorToApply
						}
					}

				}
			case "date":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				default:
					layout := "2006-01-02"
					var date_format string
					if options.Config.MetaConfig.DateFormat != "" {
						date_format = options.Config.MetaConfig.DateFormat
					} else {
						date_format = layout
					}
					str := item.Value.(string)
					t, err := time.Parse(layout, str)
					if err != nil {
						strValue = str
					} else {
						strValue = t.Format(date_format)
					}

				}
			case "datetime":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				default:
					layout := "2006-01-02 15:04:05"
					var datetime_format string
					if options.Config.MetaConfig.DatetimeFormat != "" {
						datetime_format = options.Config.MetaConfig.DatetimeFormat
					} else {
						datetime_format = layout
					}
					str := item.Value.(string)
					t, err := time.Parse(layout, str)
					if err != nil {
						strValue = str
					} else {
						strValue = t.Format(datetime_format)
					}

				}
			case "float", "monetary":
				strValue = fmt.Sprintf("%f", item.Value.(float64))
			case "integer":
				strValue = fmt.Sprintf("%d", int(item.Value.(float64)))
			case "many2one":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = "⛔"
					} else if options.Config.MetaConfig.ShowIDs {
						strValue = fmt.Sprintf("%v", item.Value)
					} else if len(item.Value.([]interface{})) == 2 {
						strValue = item.Value.([]interface{})[1].(string)
					} else {
						strValue = fmt.Sprintf("%v", item.Value)
					}
				}
			case "many2many":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = "⛔"
					} else if options.Config.MetaConfig.ShowIDs {
						strValue = fmt.Sprintf("%v", item.Value)
					} else {
						strValue = getX2ManyStrValue(options, fieldName, tableData.X2ManyData, item.Value.([]interface{}))
					}
				}
			case "one2many":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = "⛔"
					} else {
						strValue = getX2ManyStrValue(options, fieldName, tableData.X2ManyData, item.Value.([]interface{}))
					}
				}
			case "bool", "boolean":
				if item.Value.(bool) {
					strValue = emoji.Sprint("🟢")
				} else {
					strValue = emoji.Sprint("🔴")
				}
			case "binary":
				switch item.Value.(type) {
				case bool:
					strValue = "⛔"
				default:
					strValue = "📚"
				}
			case "raw":
				strValue = fmt.Sprintf("%+v", item.Value)
			default:
				log.Error("new type detected: ", item.Type, " Value=", item.Value, " HEADER=", headerValueMap[column], " model=", tableData.Model, " Type=", fmt.Sprintf("%T", item.Value))
				strValue = emoji.Sprint("👽")
			}
			tableCell.SetText(strValue)
			if column == 0 {
				tableCell.SetTextColor(options.Skin.TitleColor)
			}
			table.SetCell(row+1, column+1, tableCell)
		}
		for column := range tableData.Header {
			if column == 0 {
				continue
			}
			tableCell := table.GetCell(row+1, column)
			if fgColor != options.Skin.TableBodyFgColor {
				tableCell.SetTextColor(fgColor)
			}
		}
		idx++
	}
	title := fmt.Sprintf(" [#76b4da]([#02fffe]%s[#76b4da]) [#76b4da][[#FFFFFF]%d[#76b4da]] [#76b4da]Page [#fe00fe]<%d/%d> ", tableData.Title, tableData.Count, tableData.Page, tableData.Pages)
	table.SetBorder(true).SetTitle(title)
	return table
}

func showSearchReadResult(command *odoo.Command, options *Options) error {
	odooCfg := options.OdooCfg
	tableData, err := data.RunSearchReadCommand(odooCfg, command, options.Config)
	if err != nil {
		return err
	}
	return showData(tableData, options)
}
func showData(tableData data.Data, options *Options) error {
	table := getTableScreen(tableData, options)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = checkTableDrillDownShortcuts(table, event, options)
		event = checkTableNavigationShortcuts(table, event, options)
		event = checkTableSearchBarShortcuts(event, options)
		return event
	})
	clearMainContainer(options)
	options.MainContainer.AddItem(table, 0, 0, 1, 1, 0, 0, true)
	options.Table = table
	setTableFocus(options)
	if len(tableData.Lines) == 0 {
		return showInfo("🥺 no data found 🥺", options, tcell.ColorRed)
	}
	return nil
}
func removeTable(options *Options) {
	if mainContainerHasTable(options) {
		var emptyTable *tview.Table
		options.Table = emptyTable
	}
}
func mainContainerHasTable(options *Options) bool {
	var emptyTable *tview.Table
	if reflect.DeepEqual(options.Table, emptyTable) {
		return false
	} else {
		return true
	}
}
func getTableModelID(options *Options) (string, string, int, error) {
	if !mainContainerHasTable(options) {
		return "", "", 0, fmt.Errorf("There is no active table")
	}
	table := options.Table
	row, _ := table.GetSelection()
	var id int
	var model, name string
	for col := 0; col < table.GetColumnCount(); col++ {
		cell := table.GetCell(row, col)
		if cell.Reference == nil {
			continue
		}
		odooCellReference := cell.Reference.(*OdooCellReference)
		if odooCellReference.FieldName == "id" {
			var err error
			id, err = strconv.Atoi(cell.Text)
			if err != nil {
				return "", "", 0, err
			}
			model = odooCellReference.Model
		}
		if odooCellReference.FieldName == "name" {
			name = cell.Text
		}
	}

	if id == 0 {
		msg := "Can not found any ID from this table"
		return "", "", 0, fmt.Errorf(msg)
	}
	if name == "" {
		name = strconv.Itoa(id)
	}
	return model, name, id, nil
}
func setTableFocus(options *Options) {
	if mainContainerHasTable(options) {
		options.App.SetFocus(options.Table)
	} else {
		showInfo("No table found", options, tcell.ColorRed)
	}
}

func checkTableNavigationShortcuts(table *tview.Table, event *tcell.EventKey, options *Options) *tcell.EventKey {
	if event.Rune() == 'n' {
		goToNextPage(options)
	} else if event.Rune() == 'p' {
		goToPreviousPage(options)
	} else if event.Rune() == 'f' {
		goToFirstPage(options)
	} else if event.Rune() == 'l' {
		goToLastPage(options)
	} else if event.Rune() == 'r' {
		refreshPage(options, true)
	} else if event.Rune() == '?' {
		showHome(options)
	}
	return event
}

func checkTableSearchBarShortcuts(event *tcell.EventKey, options *Options) *tcell.EventKey {
	if event.Rune() == '/' {
		setSearchBarText("/", options)
		setSearchBarFocus(options)
	} else if event.Rune() == '@' {
		setSearchBarText("@", options)
		setSearchBarFocus(options)
	} else if event.Rune() == '#' {
		setSearchBarText("#", options)
		setSearchBarFocus(options)
	} else if event.Rune() == ':' {
		setSearchBarText(":", options)
		setSearchBarFocus(options)
	} else if event.Rune() == '!' {
		setSearchBarText("!", options)
		setSearchBarFocus(options)
	} else if event.Rune() == '>' {
		setSearchBarText(">", options)
		setSearchBarFocus(options)
	}
	return event
}
func checkTableDrillDownShortcuts(table *tview.Table, event *tcell.EventKey, options *Options) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
		model, _, id, err := getTableModelID(options)
		if err != nil {
			showInfo(err.Error(), options, tcell.ColorRed)
		} else {
			lastCommand, err := options.CommandsHistory.GetCommand()
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			rcmds, err := odoo.GetRelatedCommands(options.OdooCfg, lastCommand, model)
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
			rcmd := rcmds[0] //we have one or more
			rcmd.SetIDs([]int{id})
			cmd := rcmd.GetCommand(options.OdooCfg)
			options.CommandsHistory.AddCommand(cmd)
			err = showSearchReadResult(cmd, options)
			if err != nil {
				showInfo(err.Error(), options, tcell.ColorRed)
				return event
			}
		}
	}
	return event
}
