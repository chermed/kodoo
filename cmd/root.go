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
	Version = "v0.3.2"
	URL     = "https://github.com/chermed/kodoo"
)

var (
	rootCmd = &cobra.Command{
		Use:    "kodoo",
		Short:  "Kodoo is terminal UI for Odoo",
		Long:   `A Fast and Flexible terminal UI for Odoo (visit https://github.com/chermed/kodoo).`,
		PreRun: loadConfigAndViper,
		Run:    startUI,
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("macro", "m", "", "the name of the macro to execute")
	rootCmd.PersistentFlags().BoolP("readonly", "", false, "activate the global readonly mode")
	rootCmd.PersistentFlags().BoolP("zen-mode", "", false, "activate the zen mode")
	rootCmd.PersistentFlags().BoolP("debug", "", false, "debug mode")
	rootCmd.PersistentFlags().StringP("server", "s", "", "the odoo server to use")
	rootCmd.PersistentFlags().BoolP("no-header", "", false, "hide the header")
	rootCmd.PersistentFlags().BoolP("no-password", "", false, "hide the password")
	rootCmd.PersistentFlags().IntP("limit", "l", 0, "the default limit of records to load")
	rootCmd.PersistentFlags().StringP("logfile", "", "logfile.log", "The path to the logfile")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initConfigCmd)
	rootCmd.AddCommand(sampleConfigCmd)
}

func loadConfigAndViper(cmd *cobra.Command, args []string) {
	config.LoadConfig()
	viper.BindPFlag("config.readonly", cmd.PersistentFlags().Lookup("readonly"))
	viper.BindPFlag("config.zen_mode", cmd.PersistentFlags().Lookup("zen-mode"))
	viper.BindPFlag("config.debug", cmd.PersistentFlags().Lookup("debug"))
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
		color.Redln(fmt.Errorf("error when loading the configuration"))
		os.Exit(-1)
	}
	var filename string = viper.GetString("config.logfile")
	var debug bool = viper.GetBool("config.debug")
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	var log = logrus.New()
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}
	if err == nil {
		log.Out = f
	} else {
		log.Warn("failed to log to file, using default stderr")
	}
	cfg.Log = log
	cfg.Version = Version
	cfg.URL = URL
	if len(cfg.Servers) == 0 {
		color.Redln(fmt.Errorf("no server defined in the configuration"))
		os.Exit(-1)
	}
	if cfg.MetaConfig.ZenMode {
		cfg.MetaConfig.NoHeader = true
		cfg.MetaConfig.Refresh.Startup = true
		if cfg.MetaConfig.Refresh.IntervalSeconds <= 0 {
			cfg.MetaConfig.Refresh.IntervalSeconds = 4
		}
		if cfg.MetaConfig.DefaultMacro == "" {
			color.Redln(fmt.Errorf("you have to define the default macro"))
			os.Exit(-1)
		}
		if cfg.MetaConfig.DefaultServer == "" {
			color.Redln(fmt.Errorf("you have to define the default server"))
			os.Exit(-1)
		}
	}
	ui.AppRun(cfg)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
