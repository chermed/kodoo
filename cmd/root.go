package cmd

import (
	"fmt"
	"os"

	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/internal/ui"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Version = "v0.1.0"
	URL     = "https://github.com/chermed/kodoo"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "kodoo",
		Short: "Kodoo is terminal client for Odoo",
		Long: `A Fast and Flexible data management for Odoo
				  love by chermed and friends in Go.
				  Complete documentation is available at https://github.com/chermed/kodoo`,
		PreRun: loadConfigAndViper,
		Run:    startUI,
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("macro", "m", "", "The name of the macro to execute")
	rootCmd.PersistentFlags().StringP("server", "s", "", "The odoo server to use")
	rootCmd.PersistentFlags().BoolP("no-header", "", false, "Hide the header")
	rootCmd.PersistentFlags().BoolP("no-password", "", false, "Hide the password")
	rootCmd.PersistentFlags().IntP("limit", "l", 0, "The limit of records to load")
	rootCmd.PersistentFlags().StringP("logfile", "", "logfile.log", "The path to the logfile")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initConfigCmd)
	rootCmd.AddCommand(sampleConfigCmd)
}

func loadConfigAndViper(cmd *cobra.Command, args []string) {
	config.LoadConfig()
	viper.BindPFlag("config.default_macro", cmd.PersistentFlags().Lookup("macro"))
	viper.BindPFlag("config.default_server", cmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("config.default_limit", cmd.PersistentFlags().Lookup("limit"))
	viper.BindPFlag("config.no_header", cmd.PersistentFlags().Lookup("no-header"))
	viper.BindPFlag("config.no_password", cmd.PersistentFlags().Lookup("no-password"))
	viper.BindPFlag("config.logfile", cmd.PersistentFlags().Lookup("logfile"))
}
func startUI(cmd *cobra.Command, args []string) {
	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		color.Redln(fmt.Errorf("Error when loading the configuration"))
		os.Exit(-1)
	}
	var filename string = viper.GetString("config.logfile")
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	var log = logrus.New()

	if err == nil {
		log.Out = f
	} else {
		log.Warn("Failed to log to file, using default stderr")
	}
	cfg.Log = log
	cfg.Version = Version
	cfg.URL = URL
	ui.AppRun(cfg)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
