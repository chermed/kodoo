package odoo

import (
	mapstructure "github.com/mitchellh/mapstructure"
)

func (server *Server) ListDatabases(odooCfg *OdooConfig) ([]string, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallDB(odooCfg, "list")
	if err != nil {
		log.Error("ListDatabases CallDB: ", err)
		return []string{}, err
	}
	var odooDBsResult []string
	err = mapstructure.Decode(odooResponse.Result, &odooDBsResult)
	if err != nil {
		log.Error("ListDatabases Decode: ", err)
		return []string{}, err
	}
	return odooDBsResult, nil
}
func (server *Server) GetServerVersion(odooCfg *OdooConfig) (string, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallDB(odooCfg, "server_version")
	if err != nil {
		log.Error("GetServerVersion CallDB: ", err)
		return "", err
	}
	var odooServerVersion string
	err = mapstructure.Decode(odooResponse.Result, &odooServerVersion)
	if err != nil {
		log.Error("GetServerVersion Decode: ", err)
		return "", err
	}
	return odooServerVersion, nil
}
