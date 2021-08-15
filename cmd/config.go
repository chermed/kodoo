package cmd

import (
	"fmt"

	"github.com/chermed/kodoo/internal/config"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var initConfigCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Create a basic configuration file if not exists",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.InitConfigFiles()
		if err != nil {
			color.Redln(err)
		} else {
			color.Greenln(fmt.Sprintf("The configuration file is created to %s", path))
		}
	},
}

var sampleConfigCmd = &cobra.Command{
	Use:   "sample-config",
	Short: "Show a sample configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetSampleConfig())
	},
}
