package odoo

type RelatedCommand struct {
	Model   string
	Field   string
	IDs     []int
	Score   int
	Context OdooContext
}

func NewRelatedCommand(odooCfg *OdooConfig, model string, field string, ids []int, context OdooContext) *RelatedCommand {
	rcmd := &RelatedCommand{
		Model:   model,
		Field:   field,
		Context: context,
		IDs:     ids,
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
