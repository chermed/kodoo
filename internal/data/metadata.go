package data

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func ReadMetadata(cmd *odoo.Command, cfg *config.Config, odooCfg *odoo.OdooConfig) (odoo.OdooMetadataResult, error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return odoo.OdooMetadataResult{}, err
	}
	return server.Metadata(cmd, odooCfg)
}
