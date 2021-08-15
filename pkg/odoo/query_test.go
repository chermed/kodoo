package odoo

import (
	"reflect"
	"testing"
)

func TestQueryToCommand(t *testing.T) {
	queryStr := `    sale.order id +name -state partner_id id=3   state=draft,confirm comment=Yesthere!  name~11 `
	cmd := &Command{}
	err := StringToCommand(cmd, queryStr)
	if err != nil {
		t.Error(err)
	}
	if cmd.Model != "sale.order" {
		t.Errorf("The expected model is sale.order, found=%v", cmd.Model)
	}
	if !reflect.DeepEqual(cmd.Fields, []string{"id", "name", "state", "partner_id"}) {
		t.Errorf("The expected fields aren't good, found=%v", cmd.Fields)
	}
	if cmd.Order != "name asc, state desc" {
		t.Errorf("The expected order isn't good, found=%v", cmd.Order)
	}

}
