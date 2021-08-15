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
		log.Error(err)
		return OdooResponse{}, err
	}
	odooResponse, err := sendRequest(odooCfg, server, url, jsonValue)
	if err != nil {
		log.Error(err)
		return OdooResponse{}, err
	}
	session_id := ""
	for cookieName, cookieValue := range odooResponse.Cookies {
		if cookieName == "session_id" {
			session_id = cookieValue
			break
		}
	}
	var odooUserResult OdooUserResult
	err = mapstructure.Decode(odooResponse.Result, &odooUserResult)
	if err != nil {
		log.Error(err)
		return OdooResponse{}, err
	}
	uid := odooUserResult.UID
	if uid > 0 && session_id != "" {
		server.UID = uid
		server.SessionID = session_id
		server.ServerVersion = odooUserResult.ServerVersion
		return odooResponse, nil
	} else {
		return OdooResponse{}, fmt.Errorf("Can not connect to the server [%s] with login=%s and password=%s",
			server.Host,
			server.User,
			server.HiddenPassword)
	}
}
