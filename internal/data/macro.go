package data

import "github.com/chermed/kodoo/internal/config"

func GetMacros(config config.Config) Data {
	tableData := Data{
		Lines:  []map[string]RowItem{},
		Header: []string{"name", "description", "model"},
		Title:  "Macros",
		Count:  len(config.Macros),
		Pages:  1,
		Page:   1,
	}
	for name, values := range config.Macros {
		macroName := RowItem{
			Value: name,
			Type:  "string",
		}
		macroDescription := RowItem{
			Value: values.Description,
			Type:  "string",
		}
		macroModel := RowItem{
			Value: values.Model,
			Type:  "string",
		}
		tableData.Lines = append(tableData.Lines, map[string]RowItem{
			"name":        macroName,
			"description": macroDescription,
			"model":       macroModel,
		})
	}
	return tableData
}
