package data

import (
	"fmt"
	"sort"

	"github.com/chermed/kodoo/internal/config"
	"github.com/chermed/kodoo/pkg/odoo"
)

func GetFields(config *config.Config, odooCfg *odoo.OdooConfig, model string) (Data, error) {
	server, err := GetCurrentServer(config)
	if err != nil {
		return Data{}, err
	}
	cmd := odoo.NewCommand(
		odooCfg,
		model,
		make([][]interface{}, 0),
		[]string{},
		0,
		"",
		make(odoo.OdooContext, 0),
	)
	fields, err := server.Fields(cmd, odooCfg)
	if err != nil {
		return Data{}, err
	}
	tableData := Data{
		Lines: []map[string]RowItem{},
		Header: []string{
			"name",
			"description",
			"type",
			"relation",
			"relation_field",
			"manual",
			"required",
			"readonly",
			"selection",
			"store",
			"company_dependent",
			"searchable",
			"sortable",
		},
		Title: fmt.Sprintf("Fields of %s", model),
		Count: len(fields),
		Pages: 1,
		Page:  1,
	}
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, name := range keys {
		spec := fields[name]
		fieldName := RowItem{
			Value: name,
			Type:  "string",
		}
		fieldDescription := RowItem{
			Value: spec.Description,
			Type:  "string",
		}
		fieldType := RowItem{
			Value: spec.Type,
			Type:  "string",
		}
		fieldRelation := RowItem{
			Value: spec.Relation,
			Type:  "string",
		}
		fieldRelationField := RowItem{
			Value: spec.RelationField,
			Type:  "string",
		}
		fieldManual := RowItem{
			Value: spec.Manual,
			Type:  "bool",
		}
		fieldRequired := RowItem{
			Value: spec.Required,
			Type:  "bool",
		}
		fieldReadonly := RowItem{
			Value: spec.Readonly,
			Type:  "bool",
		}
		fieldSelection := RowItem{
			Value: spec.Selection,
			Type:  "raw",
		}
		fieldStore := RowItem{
			Value: spec.Store,
			Type:  "bool",
		}
		fieldCompanyDependent := RowItem{
			Value: spec.CompanyDependent,
			Type:  "bool",
		}
		fieldSearchable := RowItem{
			Value: spec.Searchable,
			Type:  "bool",
		}
		fieldSortable := RowItem{
			Value: spec.Sortable,
			Type:  "bool",
		}
		tableData.Lines = append(tableData.Lines, map[string]RowItem{
			"name":              fieldName,
			"description":       fieldDescription,
			"type":              fieldType,
			"relation":          fieldRelation,
			"relation_field":    fieldRelationField,
			"manual":            fieldManual,
			"required":          fieldRequired,
			"readonly":          fieldReadonly,
			"selection":         fieldSelection,
			"store":             fieldStore,
			"company_dependent": fieldCompanyDependent,
			"searchable":        fieldSearchable,
			"sortable":          fieldSortable,
		})
	}
	return tableData, nil
}
