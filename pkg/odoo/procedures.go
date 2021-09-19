package odoo

import (
	"encoding/xml"
	"fmt"
	"sync"

	mapstructure "github.com/mitchellh/mapstructure"
)

func (server *Server) SearchRead(cmd *Command, odooCfg *OdooConfig) (OdooReadResult, error) {
	log := odooCfg.Log
	if err := cmd.UpdateCommandFields(server, odooCfg); err != nil {
		return OdooReadResult{}, err
	}
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "search_read", cmd.Domain, cmd.Fields, cmd.Offset, cmd.Limit, cmd.Order)
	if err != nil {
		return OdooReadResult{}, err
	}
	var odooReadResult OdooReadResult
	err = mapstructure.Decode(odooResponse.Result, &odooReadResult)
	if err != nil {
		log.Error("SearchRead Decode: ", err)
		return OdooReadResult{}, err
	}
	return odooReadResult, nil
}
func (server *Server) ReadOneFieldX2Many(fieldName string, relation string, x2ManyData X2ManyResult, fieldIDs []int, odooCfg *OdooConfig) error {
	odooCfg.Mutex.Lock()
	nameGetCommand := &Command{
		Model: relation,
		IDS:   fieldIDs,
	}
	nameGetResult, err := server.NameGet(nameGetCommand, odooCfg)
	if err != nil {
		return err
	}
	for _, record := range nameGetResult {
		recordID := int(record[0].(float64))
		recordName := record[1].(string)
		x2ManyData[fieldName][recordID] = recordName
	}
	odooCfg.Mutex.Unlock()

	return nil
}
func (server *Server) ReadX2Many(cmd *Command, data OdooReadResult, odooCfg *OdooConfig) (X2ManyResult, error) {
	x2ManyData := X2ManyResult{}
	var wg sync.WaitGroup
	var x2ManyError error
	for _, fieldName := range cmd.Fields {
		x2ManyData[fieldName] = make(map[int]string)
	}
	for _, fieldName := range cmd.Fields {
		spec := cmd.AllFields[fieldName]
		if spec.Type == "many2many" || spec.Type == "one2many" {
			fieldIDs := []int{}
			for _, line := range data {
				ids := line[fieldName].([]interface{})
				for _, id := range ids {
					IDFound := false
					currentID := int(id.(float64))
					for _, fieldID := range fieldIDs {
						if fieldID == currentID {
							IDFound = true
							break
						}
					}
					if !IDFound {
						fieldIDs = append(fieldIDs, currentID)
					}
				}
			}
			wg.Add(1)
			fName := fieldName
			specRelation := spec.Relation
			go func() {
				defer wg.Done()
				oneFieldErr := server.ReadOneFieldX2Many(fName, specRelation, x2ManyData, fieldIDs, odooCfg)
				if oneFieldErr != nil {
					x2ManyError = oneFieldErr
				}
			}()
		}
	}
	wg.Wait()
	return x2ManyData, x2ManyError
}

func (server *Server) Search(cmd *Command, odooCfg *OdooConfig) (OdooSearchResult, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "search", cmd.Domain, cmd.Offset, cmd.Limit, cmd.Order)
	if err != nil {
		return OdooSearchResult{}, err
	}
	var odooSearchResult OdooSearchResult
	err = mapstructure.Decode(odooResponse.Result, &odooSearchResult)
	if err != nil {
		log.Error("Search Decode: ", err)
		return OdooSearchResult{}, err
	}
	return odooSearchResult, nil
}
func (server *Server) Read(cmd *Command, ids []int, odooCfg *OdooConfig) (OdooReadResult, error) {
	log := odooCfg.Log
	if err := cmd.UpdateCommandFields(server, odooCfg); err != nil {
		return OdooReadResult{}, err
	}
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "read", ids, cmd.Fields)
	if err != nil {
		return OdooReadResult{}, err
	}
	var odooReadResult OdooReadResult
	err = mapstructure.Decode(odooResponse.Result, &odooReadResult)
	if err != nil {
		log.Error("Read Decode: ", err)
		return OdooReadResult{}, err
	}
	return odooReadResult, nil
}
func (server *Server) Count(cmd *Command, odooCfg *OdooConfig) (int, error) {
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "search_count", cmd.Domain)
	if err != nil {
		return 0, err
	}
	return int(odooResponse.Result.(float64)), nil
}
func (server *Server) Fields(cmd *Command, odooCfg *OdooConfig) (OdooFieldsResult, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "fields_get")
	if err != nil {
		return OdooFieldsResult{}, err
	}
	var odooFieldsResult OdooFieldsResult
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &odooFieldsResult,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		log.Error("Fields NewDecoder: ", err)
		return OdooFieldsResult{}, err
	}
	err = decoder.Decode(odooResponse.Result)
	if err != nil {
		log.Error("Fields Decode: ", err)
		return OdooFieldsResult{}, err
	}
	return odooFieldsResult, nil
}
func (server *Server) NameGet(cmd *Command, odooCfg *OdooConfig) (OdooNameGetResult, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "name_get", cmd.IDS)
	if err != nil {
		return OdooNameGetResult{}, err
	}
	var odooNameGetResult OdooNameGetResult
	err = mapstructure.Decode(odooResponse.Result, &odooNameGetResult)
	if err != nil {
		log.Error("NameGet Decode: ", err)
		return OdooNameGetResult{}, err
	}
	return odooNameGetResult, nil
}
func (server *Server) Metadata(cmd *Command, odooCfg *OdooConfig) (OdooMetadataResult, error) {
	log := odooCfg.Log
	if len(cmd.IDS) == 0 {
		err := fmt.Errorf("no IDS provided to read metadata")
		log.Error("Metadata: ", err)
		return OdooMetadataResult{}, err
	}
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "get_metadata", cmd.IDS)
	if err != nil {
		return OdooMetadataResult{}, err
	}
	var odooMetadataResult OdooMetadataResult
	err = mapstructure.Decode(odooResponse.Result, &odooMetadataResult)
	if err != nil {
		log.Error("Metadata Decode: ", err)
		return OdooMetadataResult{}, err
	}
	return odooMetadataResult, nil
}
func (server *Server) FieldsViewGet(cmd *Command, odooCfg *OdooConfig) ([]string, error) {
	log := odooCfg.Log
	odooResponse, err := server.CallObject(odooCfg, cmd.Model, "fields_view_get", nil, "tree")
	if err != nil {
		return []string{}, err
	}
	var odooFieldsViewGetResult OdooFieldsViewGetResult
	err = mapstructure.Decode(odooResponse.Result, &odooFieldsViewGetResult)
	if err != nil {
		log.Error("FieldsViewGet Decode: ", err)
		return []string{}, err
	}
	var odooTree OdooTree
	err = xml.Unmarshal([]byte(odooFieldsViewGetResult.Arch), &odooTree)
	if err != nil {
		log.Error("FieldsViewGet Unmarshal: ", err)
		return []string{}, err
	}
	fields := []string{}
	for _, fieldStruct := range odooTree.Field {
		if fieldStruct.Invisible == "1" {
			continue
		}
		found := false
		for _, f := range fields {
			if f == fieldStruct.Name {
				found = true
				break
			}
		}
		if !found {
			fields = append(fields, fieldStruct.Name)
		}
	}
	if len(fields) == 0 {
		err = fmt.Errorf("no field found using <fields view get>")
		log.Error("FieldsViewGet", err)
		return []string{}, err
	} else {
		fields = append(fields[:1], fields[0:]...)
		fields[0] = "id"
	}
	return fields, nil
}
func (server *Server) FieldsTree(cmd *Command, odooCfg *OdooConfig) ([]string, error) {
	log := odooCfg.Log
	treeCommand := NewCommand(odooCfg, "ir.ui.view", [][]interface{}{
		{"model", "=", cmd.Model},
		{"type", "=", "tree"},
		{"mode", "=", "primary"},
	}, []string{"arch_base"}, 1, "", OdooContext{})
	treeCommand.FieldsUpdated = true
	data, err := server.SearchRead(treeCommand, odooCfg)
	if err != nil {
		return []string{}, err
	}
	if len(data) == 0 {
		errMsg := fmt.Sprintf("There is no tree view for the model %v", cmd.Model)
		log.Error(errMsg)
		return []string{}, fmt.Errorf(errMsg)
	}
	var modelXmlValue string
	for _, line := range data {
		for key, value := range line {
			if key == "arch_base" {
				modelXmlValue = value.(string)
			}
		}
	}
	var odooTree OdooTree
	err = xml.Unmarshal([]byte(modelXmlValue), &odooTree)
	if err != nil {
		log.Error("FieldsTree Unmarshal: ", err)
		return []string{}, err
	}
	fields := []string{}
	for _, fieldStruct := range odooTree.Field {
		if fieldStruct.Invisible == "1" {
			continue
		}
		found := false
		for _, f := range fields {
			if f == fieldStruct.Name {
				found = true
				break
			}
		}
		if !found {
			fields = append(fields, fieldStruct.Name)
		}
	}
	return fields, nil
}
