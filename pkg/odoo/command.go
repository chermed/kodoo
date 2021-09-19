package odoo

import (
	"math"
	"sort"
	"sync"

	"github.com/sirupsen/logrus"
)

type OdooConfig struct {
	Log          *logrus.Logger
	DefaultLimit int
	Timeout      int
	Mutex        sync.RWMutex
	NoAutoId     bool
}
type Command struct {
	Description   string
	Model         string
	Domain        [][]interface{}
	Limit         int
	Offset        int
	Fields        []string
	IDS           []int
	AllFields     OdooFieldsResult
	FieldsUpdated bool
	Order         string
	Page          int
	Pages         int
	Count         int
	Context       OdooContext
}

func NewCommand(odooCfg *OdooConfig, model string, domain [][]interface{}, fields []string, limit int, order string, context OdooContext) *Command {
	if len(domain) == 0 {
		domain = make([][]interface{}, 0)
	}
	if limit == 0 && odooCfg.DefaultLimit > 0 {
		limit = odooCfg.DefaultLimit
	}
	cmd := &Command{
		Model:   model,
		Domain:  domain,
		Fields:  fields,
		Limit:   limit,
		Order:   order,
		Offset:  0,
		Page:    1,
		Pages:   0,
		Context: context,
	}
	return cmd
}

func NewCommandIDs(model string, ids []int) *Command {
	cmd := &Command{
		Model: model,
		IDS:   ids,
		Domain: [][]interface{}{
			{
				"id", "in", ids,
			},
		},
	}
	return cmd
}

func (cmd *Command) UseAllFields() {
	fields := []string{}
	for field := range cmd.AllFields {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	cmd.Fields = fields
}

func (cmd *Command) SetFieldsUpdated() {
	cmd.FieldsUpdated = true
}

func (cmd *Command) AreFieldsUpdated() bool {
	return cmd.FieldsUpdated
}

func (cmd *Command) SetID(id int) {
	cmd.Domain = [][]interface{}{
		{"id", "=", id},
	}
}

func (cmd *Command) UpdateCommandFields(server *Server, odooCfg *OdooConfig) error {
	log := odooCfg.Log
	if cmd.AreFieldsUpdated() {
		return nil
	}
	var wg sync.WaitGroup
	var sharedError error
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, err := server.Count(cmd, odooCfg)
		if err != nil {
			sharedError = err
		} else {
			cmd.Count = count
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		allFields, err := server.Fields(cmd, odooCfg)
		if err != nil {
			sharedError = err
		} else {
			cmd.AllFields = allFields
		}
	}()
	wg.Wait()

	if sharedError != nil {
		log.Error("UpdateCommandFields: ", sharedError)
		return sharedError
	}
	if cmd.Limit > 0 {
		cmd.Pages = int(math.Ceil(float64(cmd.Count) / float64(cmd.Limit)))
	} else {
		cmd.Pages = 1
	}
	if len(cmd.Fields) == 0 {
		fieldsViewGet, err := server.FieldsViewGet(cmd, odooCfg)
		if err == nil {
			cmd.Fields = fieldsViewGet
		}
	}
	if len(cmd.Fields) == 0 {
		fieldsTree, _ := server.FieldsTree(cmd, odooCfg)
		if len(fieldsTree) > 0 {
			cmd.Fields = fieldsTree
		}
	}
	if len(cmd.Fields) == 0 {
		for fieldName := range cmd.AllFields {
			if fieldName == "name" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "create_uid" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "create_date" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "write_uid" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "write_date" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "partner_id" {
				cmd.Fields = append(cmd.Fields, fieldName)
			} else if fieldName == "date" {
				cmd.Fields = append(cmd.Fields, fieldName)
			}
		}
	}
	idFound := false
	for _, fieldName := range cmd.Fields {
		if fieldName == "id" {
			idFound = true
		}
	}
	for fieldName := range cmd.AllFields {
		if fieldName == "id" && !idFound && !odooCfg.NoAutoId {
			cmd.Fields = append(cmd.Fields[:1], cmd.Fields[0:]...)
			cmd.Fields[0] = fieldName
		}
	}

	cmd.SetFieldsUpdated()
	return nil
}
