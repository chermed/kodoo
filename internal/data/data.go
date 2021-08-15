package data

import "github.com/chermed/kodoo/pkg/odoo"

type RowItem struct {
	Value interface{}
	Type  string
}

type Data struct {
	Lines      []map[string]RowItem
	Header     []string
	Title      string
	Count      int
	Page       int
	Pages      int
	Model      string
	X2ManyData odoo.X2ManyResult
}
