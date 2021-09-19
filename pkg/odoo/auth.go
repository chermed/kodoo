package odoo

import (
	"encoding/json"
	"fmt"

	mapstructure "github.com/mitchellh/mapstructure"
)

func (server *Server) Authenticate(odooCfg *OdooConfig) (OdooResponse, error) {
	log := odooCfg.Log
	url := cleanHost(server.Host) + "/web/session/authenticate"
	values := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "authenticate",
		"params": map[string]string{
			"db":       server.Database,
			"login":    server.User,
			"password": server.Password,
		},
	}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		log.Error("Authenticate Marshal: ", err)
		return OdooResponse{}, err
	}
	odooResponse, err := sendRequest(odooCfg, server, url, jsonValue)
	if err != nil {
		log.Error("Authenticate sendRequest: ", err)
		return OdooResponse{}, err
	}
	var odooUserResult OdooUserResult
	err = mapstructure.Decode(odooResponse.Result, &odooUserResult)
	if err != nil {
		log.Error("Authenticate Decode: ", err)
		return OdooResponse{}, err
	}
	uid := odooUserResult.UID
	if uid > 0 {
		server.UID = uid
		return odooResponse, nil
	} else {
		err = fmt.Errorf("can not connect to the server [%s] with login=%s and password=%s",
			server.Host,
			server.User,
			server.HiddenPassword)
		log.Error("Authenticate UID: ", err)
		return OdooResponse{}, err
	}
}
