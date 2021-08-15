package data

import (
	"fmt"

	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func GetServers(config config.Config) Data {
	tableHeader := []string{"name", "host", "database", "user", "password"}
	tableData := Data{
		Lines:  []map[string]RowItem{},
		Header: tableHeader,
		Title:  "Servers",
		Count:  len(config.Servers),
		Page:   1,
		Pages:  1,
	}
	for name, values := range config.Servers {
		serverName := RowItem{
			Value: name,
			Type:  "string",
		}
		serverHost := RowItem{
			Value: values.Host,
			Type:  "string",
		}
		serverDB := RowItem{
			Value: values.Database,
			Type:  "string",
		}
		serverUser := RowItem{
			Value: values.User,
			Type:  "string",
		}
		var serverPassword RowItem
		if config.MetaConfig.NoPassword {
			serverPassword = RowItem{
				Value: "*****",
				Type:  "string",
			}
		} else {
			serverPassword = RowItem{
				Value: values.Password,
				Type:  "string",
			}
		}
		tableData.Lines = append(tableData.Lines, map[string]RowItem{
			"name":     serverName,
			"host":     serverHost,
			"database": serverDB,
			"user":     serverUser,
			"password": serverPassword,
		})
	}
	return tableData
}

func GetCurrentServer(cfg *config.Config) (*odoo.Server, error) {
	if cfg.MetaConfig.DefaultServer == "" {
		for name := range cfg.Servers {
			cfg.MetaConfig.DefaultServer = name
			break
		}
	}
	if cfg.MetaConfig.DefaultServer == "" {
		return &odoo.Server{}, fmt.Errorf("No server found as default")
	}
	if currentServer, found := cfg.Servers[cfg.MetaConfig.DefaultServer]; found {
		return currentServer, nil
	} else {
		return &odoo.Server{}, fmt.Errorf("The server [%s] is not found", cfg.MetaConfig.DefaultServer)
	}
}

func SetCurrentServer(config *config.Config, name string) {
	config.MetaConfig.DefaultServer = name
}
