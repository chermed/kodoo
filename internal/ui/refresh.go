package ui

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func autoRefreshIsEnabled(options *Options) bool {
	return options.Config.MetaConfig.Refresh.Startup
}

func toggleAutoRefresh(options *Options) {
	options.Config.MetaConfig.Refresh.Startup = !options.Config.MetaConfig.Refresh.Startup
	if autoRefreshIsEnabled(options) {
		showInfo("the autorefresh is enabled", options, tcell.ColorGreen)
	} else {
		showInfo("the autorefresh is disabled", options, tcell.ColorOrange)
	}
}
func startRefreshTicker(options *Options) {
	intervalSeconds := options.Config.MetaConfig.Refresh.IntervalSeconds
	if intervalSeconds <= 0 {
		intervalSeconds = 4
	}
	ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
	go func() {
	loop:
		for {
			select {
			case <-ticker.C:
				if autoRefreshIsEnabled(options) && options.CommandsHistory.HasCommand() {
					refreshPage(options, true)
					setSearchBarFocus(options)
				}
			case <-options.Shutdown:
				break loop
			}
		}
	}()
}
