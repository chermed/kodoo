package odoo

type RelatedCommand struct {
	Model       string
	Field       string
	IDs         []int
	Score       int
	Description string
	Context     OdooContext
	OriginField string
	Type        string
}

func NewRelatedCommand(odooCfg *OdooConfig, model string, field string, ids []int, description string, ttype string, context OdooContext) *RelatedCommand {
	rcmd := &RelatedCommand{
		Model:       model,
		Field:       field,
		Context:     context,
		IDs:         ids,
		Description: description,
		Type:        ttype,
	}
	return rcmd
}

func (rcmd *RelatedCommand) SetIDs(ids []int) {
	rcmd.IDs = ids
}

func (rcmd *RelatedCommand) GetCommand(odooCfg *OdooConfig) *Command {
	cmd := NewCommand(
		odooCfg,
		rcmd.Model,
		[][]interface{}{
			{rcmd.Field, "in", rcmd.IDs},
		},
		[]string{},
		0,
		"",
		rcmd.Context,
	)
	return cmd
}
