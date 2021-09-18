package odoo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func sendRequest(odooCfg *OdooConfig, server *Server, url string, payload []byte) (OdooResponse, error) {
	log := odooCfg.Log
	timeout := odooCfg.Timeout
	if timeout <= 0 {
		timeout = 10
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	log.Debug(fmt.Sprintf("sending request to %s with payload : %s", server.GetName(), string(payload)))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Error(err)
		return OdooResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return OdooResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error("status code is: ", resp.StatusCode)
		return OdooResponse{}, fmt.Errorf("the connection to the server %s is refused", server.Host)
	}
	odooResponse := OdooResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(body)
		return OdooResponse{}, err
	}
	if err := json.Unmarshal(body, &odooResponse); err != nil {
		log.Error(err)
		return OdooResponse{}, err
	}
	if !reflect.DeepEqual(OdooError{}, odooResponse.Error) {
		return OdooResponse{}, fmt.Errorf("%v: %v",
			odooResponse.Error.Message,
			odooResponse.Error.Data.Message)
	}
	odooResponse.Cookies = make(map[string]string)
	for _, cookie := range resp.Cookies() {
		odooResponse.Cookies[cookie.Name] = cookie.Value
	}
	return odooResponse, nil

}

func (server *Server) CallObject(odooCfg *OdooConfig, object string, method string, args ...interface{}) (OdooResponse, error) {
	log := odooCfg.Log
	if server.UID == 0 {
		log.Info("need authentication")
		if _, err := server.Authenticate(odooCfg); err != nil {
			return OdooResponse{}, err
		}

	}
	url := cleanHost(server.Host) + "/jsonrpc"
	arguments := []interface{}{
		server.Database,
		strconv.Itoa(server.UID),
		server.Password,
		object,
		method,
	}
	arguments = append(arguments, args...)
	values := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "call",
		"params": map[string]interface{}{
			"method":  "execute",
			"service": "object",
			"args":    arguments,
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
	return odooResponse, nil
}
func (server *Server) CallDB(odooCfg *OdooConfig, method string) (OdooResponse, error) {
	log := odooCfg.Log
	url := cleanHost(server.Host) + "/jsonrpc"
	arguments := []interface{}{}
	values := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "call",
		"params": map[string]interface{}{
			"method":  method,
			"service": "db",
			"args":    arguments,
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
	return odooResponse, nil
}
