package data

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func ReadMetadata(cfg *config.Config, odooCfg *odoo.OdooConfig, commandsHistory *CommandsHistory) (odoo.OdooMetadataResult, error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return odoo.OdooMetadataResult{}, err
	}
	cmd, err := commandsHistory.GetCommand()
	if err != nil {
		return odoo.OdooMetadataResult{}, err
	}
	return server.Metadata(cmd, odooCfg)
}
