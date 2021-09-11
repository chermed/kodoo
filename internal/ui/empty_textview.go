package ui

import "github.com/rivo/tview"

func getEmptyTextView(options *Options) *tview.TextView {
	emptyTextView := tview.NewTextView()
	emptyTextView.SetText(" ")
	emptyTextView.SetBackgroundColor(options.Skin.BackgroundColor)
	return emptyTextView
}
