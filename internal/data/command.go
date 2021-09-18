package data

import (
	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func RunRemoteFunction(odooCfg *odoo.OdooConfig, cfg *config.Config, object string, method string, args ...interface{}) (odoo.OdooResponse, error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return odoo.OdooResponse{}, err
	}
	data, err := server.CallObject(odooCfg, object, method, args...)
	if err != nil {
		return odoo.OdooResponse{}, err
	}
	return data, nil
}

func RunSearchReadCommand(odooCfg *odoo.OdooConfig, command *odoo.Command, cfg *config.Config) (Data, error) {
	server, err := GetCurrentServer(cfg)
	if err != nil {
		return Data{}, err
	}
	data, err := server.SearchRead(command, odooCfg)
	if err != nil {
		return Data{}, err
	}
	x2ManyData, err := server.ReadX2Many(command, data, odooCfg)
	if err != nil {
		return Data{}, err
	}
	title := command.Model
	if command.Description != "" {
		title = command.Description
	}
	tableData := Data{
		Lines:      []map[string]RowItem{},
		Title:      title,
		Header:     command.Fields,
		Count:      command.Count,
		Page:       command.Page,
		Pages:      command.Pages,
		Model:      command.Model,
		X2ManyData: x2ManyData,
	}
	for _, line := range data {
		rowItems := make(map[string]RowItem)
		for key, value := range line {
			rowItems[key] = RowItem{
				Value: value,
				Type:  command.AllFields[key].Type,
			}
		}
		tableData.Lines = append(tableData.Lines, rowItems)
	}
	return tableData, nil
}

func GetRelatedCommands(commandHistory *CommandsHistory, odooCfg *odoo.OdooConfig) (rcmds []odoo.RelatedCommand, err error) {
	lastCommand, err := commandHistory.GetCommand()
	if err != nil {
		return rcmds, err
	}
	rcmds, err = odoo.GetRelatedCommands(odooCfg, lastCommand)
	if err != nil {
		return rcmds, err
	}
	return rcmds, err
}
