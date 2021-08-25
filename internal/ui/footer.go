package ui

import (
	"github.com/rivo/tview"
)

func clearFooter(options *Options) {
	emptyTextView := getEmptyTextView(options)
	options.Footer.AddItem(emptyTextView, 0, 0, 1, 1, 0, 0, false)
	options.Footer.Clear()

}
func setupFooter(grid *tview.Grid, options *Options) *tview.Grid {
	footer := grid.SetRows(1).SetColumns(0)
	footer.SetBackgroundColor(options.Skin.BackgroundColor)
	footer.AddItem(getEmptyTextView(options), 0, 0, 1, 1, 0, 0, false)
	return footer
}
