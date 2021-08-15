package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chermed/kodoo/pkg/odoo"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	sampleConfig = `servers:
  local:
    host: http://localhost:8069/
    database: demo
    user: admin
    password: admin
config:
  default_server: local
  default_limit: 30
  default_macro: customers
  timeout: 4
  refresh:
    startup: false
    interval_seconds: 10
  show_ids: false
  date_format: 02 Jan 06
  datetime_format: 02 Jan 06 15:04 MST
  no_header: false
  no_password: false
macros:
  products:
    model: product.product
  customers:
    description: customers
    model: res.partner
    domain:
    - ["customer_rank", ">", 0]
  suppliers:
    description: suppliers
    model: res.partner
    domain:
    - ["supplier_rank", ">", 0]
    fields:
    - id
    - name
    - street
    - street2
    - zip
    - country_id
    order: name asc
    limit: 4`
)

type Macro struct {
	Description, Model string
	Limit              int
	Domain             [][]interface{}
	Fields             []string
	Order              string
	Context            odoo.OdooContext
}
type RefreshConfig struct {
	Startup         bool `mapstructure:"startup"`
	IntervalSeconds int  `mapstructure:"interval_seconds"`
}
type MainConfig struct {
	DefaultServer  string        `mapstructure:"default_server"`
	DefaultLimit   int           `mapstructure:"default_limit"`
	DefaultMacro   string        `mapstructure:"default_macro"`
	Refresh        RefreshConfig `mapstructure:"refresh"`
	ShowIDs        bool          `mapstructure:"show_ids"`
	DateFormat     string        `mapstructure:"date_format"`
	DatetimeFormat string        `mapstructure:"datetime_format"`
	NoHeader       bool          `mapstructure:"no_header"`
	NoPassword     bool          `mapstructure:"no_password"`
	Timeout        int           `mapstructure:"timeout"`
}
type MapInfo struct {
	Relation string `mapstructure:"relation"`
	Field    string `mapstructure:"field"`
}
type MapConfig map[string][]MapInfo
type Config struct {
	Servers    map[string]*odoo.Server `mapstructure:"servers"`
	MetaConfig *MainConfig             `mapstructure:"config"`
	Macros     map[string]Macro        `mapstructure:"macros"`
	Maps       MapConfig               `mapstructure:"maps"`
	Log        *logrus.Logger
	URL        string
	Version    string
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	return homeDir
}
func getKodooFolder() string {
	return filepath.Join(getHomeDir(), ".kodoo")
}
func getConfigFilePath() string {
	return filepath.Join(getKodooFolder(), "config.yaml")
}
func InitConfigFiles() (string, error) {
	kodooFolder := getKodooFolder()
	_, err := os.Stat(kodooFolder)
	if os.IsNotExist(err) {
		err := os.MkdirAll(kodooFolder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	path := getConfigFilePath()
	if fileExists(path) {
		return "", fmt.Errorf("The configuration file already exists")
	}
	err = ioutil.WriteFile(path, []byte(sampleConfig), 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}
func GetSampleConfig() string {
	return sampleConfig
}
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(getKodooFolder())
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		color.Redln(fmt.Errorf("Fatal error: %w \n", err))
		os.Exit(255)
	}
}
