package ui

import (
	"fmt"

	"github.com/chermed/kodoo/internal/data"
)

func checkReadonly(options *Options) (err error) {
	currentServer, err := data.GetCurrentServer(options.Config)
	if err != nil {
		return err
	}
	if options.Config.MetaConfig.Readonly {
		return fmt.Errorf("global readonly mode is enabled")
	}
	if currentServer.Readonly {
		return fmt.Errorf("this database is on mode readonly")
	}
	return err
}
