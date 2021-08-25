package ui

import (
	"github.com/rivo/tview"
)

func addField(options *Options, grid *tview.Grid, row int, col int, label string, value string, isError bool, margin bool) (*tview.TextView, *tview.TextView, *tview.TextView) {
	if margin {
		label = " " + label
	}
	lbl := tview.NewTextView().SetText(label)
	lbl.SetBackgroundColor(options.Skin.BackgroundColor)
	val := tview.NewTextView().SetText(value)
	val.SetBackgroundColor(options.Skin.BackgroundColor)
	colon := tview.NewTextView()
	colon.SetBackgroundColor(options.Skin.BackgroundColor)
	if label == "" && value == "" {
		colon.SetText("")
	} else {
		colon.SetText(":")
	}
	if isError {
		lbl.SetTextColor(options.Skin.FgErrorColor)
		colon.SetTextColor(options.Skin.FgErrorColor)
		val.SetTextColor(options.Skin.FgErrorColor)
	} else {
		lbl.SetTextColor(options.Skin.FieldLabelColor)
		colon.SetTextColor(options.Skin.FieldColonColor)
		val.SetTextColor(options.Skin.FieldValueColor)
	}
	grid.AddItem(lbl, row, col, 1, 1, 0, 0, false)
	grid.AddItem(colon, row, col+1, 1, 1, 0, 0, false)
	grid.AddItem(val, row, col+2, 1, 1, 0, 0, false)
	return lbl, colon, val
}
