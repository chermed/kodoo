package data

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func ListDatabases(cfg *config.Config, odooCfg *odoo.OdooConfig) (databases []string, err error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return databases, err
	}
	return server.ListDatabases(odooCfg)
}
