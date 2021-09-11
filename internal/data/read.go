package data

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func Read(cmd *odoo.Command, ids []int, cfg *config.Config, odooCfg *odoo.OdooConfig) (odoo.OdooReadResult, error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return odoo.OdooReadResult{}, err
	}
	return server.Read(cmd, ids, odooCfg)
}
