package ui

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/chermed/kodoo/internal/data"
	"github.com/chermed/kodoo/pkg/kotils"
	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/kyokomi/emoji"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	SelectionFieldName = "__selection field__"
)

type OdooCellReference struct {
	FieldName string
	Model     string
	ID        int
}

func getX2ManyStrValue(options *Options, fieldName string, x2ManyData odoo.X2ManyResult, ids []interface{}) string {
	names := []string{}
	for _, id := range ids {
		names = append(names, x2ManyData[fieldName][int(id.(float64))])
	}
	return strings.Join(names, ", ")
}
func getTableScreen(tableData data.Data, options *Options) *tview.Table {
	zenMode := options.Config.MetaConfig.ZenMode
	table := tview.NewTable()
	if !zenMode {
		table.SetFixed(1, 2)
		table.SetSelectable(true, false)
	}
	table.SetBackgroundColor(options.Skin.BackgroundColor)
	table.SetTitleColor(options.Skin.TitleColor)
	if options.Config.MetaConfig.ZenMode {
		table.SetBorders(true)
		table.SetBordersColor(options.Skin.BorderColor)
	}
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
		headerAliasValue := data.GetAliasValue(*options.Config, headerValue, tableData.Model)
		headerDisplayedName := strings.ToUpper(headerAliasValue.Value)
		if zenMode {
			headerDisplayedName = "[::b]" + headerDisplayedName
		}
		tableCell := tview.NewTableCell(headerDisplayedName).
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
		lineID := 0
		fgColor := options.Skin.TableBodyFgColor
		selectionBodyTableCell := tview.NewTableCell("").
			SetAlign(tview.AlignCenter).
			SetSelectable(true).
			SetExpansion(0).
			SetBackgroundColor(options.Skin.BackgroundColor).
			SetTextColor(options.Skin.TitleColor)
		if kotils.IntInSlice(idx, tableData.Selection) {
			selectionBodyTableCell.SetText("*")
			table.Select(row+1, 0)
		}
		selectionCellReference := &OdooCellReference{
			Model:     tableData.Model,
			FieldName: SelectionFieldName,
		}
		selectionBodyTableCell.SetReference(selectionCellReference)
		table.SetCell(row+1, 0, selectionBodyTableCell)
		for column := range tableData.Header {
			tableCell := tview.NewTableCell("").
				SetAlign(tview.AlignLeft).
				SetSelectable(true).
				SetExpansion(1).
				SetBackgroundColor(options.Skin.BackgroundColor).
				SetTextColor(fgColor)
			fieldName := headerValueMap[column]
			odooCellReference := &OdooCellReference{
				FieldName: fieldName,
				Model:     tableData.Model,
			}
			tableCell = tableCell.SetReference(odooCellReference)
			item := lines[fieldName]
			strValue := ""
			contentAlign := tview.AlignLeft
			switch item.Type {
			case "char", "string", "selection", "text":
				switch item.Value.(type) {
				case bool:
					strValue = emoji.Sprint("???")
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
					strValue = emoji.Sprint("???")
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
					strValue = emoji.Sprint("???")
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
				contentAlign = tview.AlignRight
				if fieldName == "id" {
					lineID = int(item.Value.(float64))
				}
				strValue = fmt.Sprintf("%f", item.Value.(float64))
			case "integer":
				contentAlign = tview.AlignRight
				if fieldName == "id" {
					lineID = int(item.Value.(float64))
				}
				strValue = fmt.Sprintf("%d", int(item.Value.(float64)))
			case "many2one":
				switch item.Value.(type) {
				case bool:
					strValue = emoji.Sprint("???")
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = emoji.Sprint("???")
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
					strValue = emoji.Sprint("???")
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = emoji.Sprint("???")
					} else if options.Config.MetaConfig.ShowIDs {
						strValue = fmt.Sprintf("%v", item.Value)
					} else {
						strValue = getX2ManyStrValue(options, fieldName, tableData.X2ManyData, item.Value.([]interface{}))
					}
				}
			case "one2many":
				switch item.Value.(type) {
				case bool:
					strValue = emoji.Sprint("???")
				case []interface{}:
					if len(item.Value.([]interface{})) == 0 {
						strValue = emoji.Sprint("???")
					} else {
						strValue = getX2ManyStrValue(options, fieldName, tableData.X2ManyData, item.Value.([]interface{}))
					}
				}
			case "bool", "boolean":
				if item.Value.(bool) {
					strValue = emoji.Sprint("????")
				} else {
					strValue = emoji.Sprint("????")
				}
			case "binary":
				switch item.Value.(type) {
				case bool:
					strValue = emoji.Sprint("???")
				default:
					strValue = emoji.Sprint("????")
				}
			case "raw":
				strValue = fmt.Sprintf("%+v", item.Value)
			default:
				strValue = fmt.Sprintf("%+v", item.Value)
			}
			aliasValue := data.GetAliasValue(*options.Config, strValue, tableData.Model)
			if aliasValue.Color != "" {
				tableCell.SetTextColor(tcell.GetColor(aliasValue.Color).TrueColor())
			}
			if zenMode {
				tableCell.SetText("[::b]" + aliasValue.Value)
			} else {
				tableCell.SetText(aliasValue.Value)
			}
			tableCell.SetAlign(contentAlign)
			if column == 0 && !zenMode {
				tableCell.SetTextColor(options.Skin.TitleColor)
			}
			table.SetCell(row+1, column+1, tableCell)
		}
		for column := range tableData.Header {
			tableCell := table.GetCell(row+1, column)
			cellRef := tableCell.GetReference().(*OdooCellReference)
			cellRef.ID = lineID
			if column == 0 {
				continue
			}
			if fgColor != options.Skin.TableBodyFgColor {
				tableCell.SetTextColor(fgColor)
			}
		}
		idx++
	}
	title := fmt.Sprintf(" [#02fffe::b]%s[#76b4da:bl] [#76b4da][[#FFFFFF]%d[#76b4da]] [#76b4da]PAGE [#fe00fe]<%d/%d> ", tableData.Title, tableData.Count, tableData.Page, tableData.Pages)
	table.SetTitle(title)
	table.SetBorder(true)
	if zenMode {
		table.SetBorderPadding(2, 0, 0, 0)
		table.RemoveColumn(0)
		table.SetBorderColor(options.Skin.BackgroundColor)
	}
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
		event = checkTableSelection(table, event, options)
		return event
	})
	clearMainContainer(options)
	options.MainContainer.AddItem(table, 0, 0, 1, 1, 0, 0, true)
	options.Table = table
	setTableFocus(options)
	if len(tableData.Lines) == 0 {
		return showInfo("???? no data found ????", options, tcell.ColorRed)
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
func getTableModelIDs(options *Options, selection bool) (model string, ids []int, err error) {
	if !mainContainerHasTable(options) {
		return model, ids, fmt.Errorf("there is no active table")
	}
	table := options.Table
	if selection {
		row, _ := table.GetSelection()
		cell := table.GetCell(row, 0)
		if cell.Reference != nil {
			odooCellReference := cell.Reference.(*OdooCellReference)
			model = odooCellReference.Model
			ids = append(ids, odooCellReference.ID)
		}
	} else {
		for row := 0; row < table.GetRowCount(); row++ {
			for col := 0; col < table.GetColumnCount(); col++ {
				cell := table.GetCell(row, 0)
				if cell.Reference == nil {
					continue
				}
				odooCellReference := cell.Reference.(*OdooCellReference)
				if odooCellReference.FieldName == SelectionFieldName && cell.Text != "" {
					model = odooCellReference.Model
					ids = append(ids, odooCellReference.ID)
					break
				}
			}
		}
	}

	if len(ids) == 0 {
		err = fmt.Errorf("please select some items before this action")
	}
	return model, ids, err
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
		goToNextPage(options, false)
	} else if event.Rune() == 'p' {
		goToPreviousPage(options, false)
	} else if event.Rune() == '^' {
		goToFirstPage(options, false)
	} else if event.Rune() == '$' {
		goToLastPage(options, false)
	} else if event.Rune() == 'r' {
		refreshPage(options, true)
	} else if event.Rune() == 'd' {
		showDetails(options)
	} else if event.Rune() == 'm' {
		showMetadata(options)
	} else if event.Rune() == 's' {
		showShortcuts(options)
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
		_, ids, err := getTableModelIDs(options, true)
		if err != nil {
			showInfo(err.Error(), options, tcell.ColorRed)
		} else {
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
			rcmd := rcmds[0] //we have one or more
			rcmd.SetIDs(ids)
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
func checkTableSelection(table *tview.Table, event *tcell.EventKey, options *Options) *tcell.EventKey {
	if event.Rune() == ' ' {
		row, _ := table.GetSelection()
		selectionTableCell := table.GetCell(row, 0)
		if selectionTableCell.Text == "" {
			selectionTableCell.SetText("*")
		} else {
			selectionTableCell.SetText("")
		}
	}
	return event
}
