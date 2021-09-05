package odoo

import (
	"fmt"
	"sort"
	"strings"

	strLib "github.com/mgutz/str"
)

func GetRelatedCommands(odoocfg *OdooConfig, lastCommand *Command) ([]RelatedCommand, error) {
	rcmds := []RelatedCommand{}
	for fieldName, spec := range lastCommand.AllFields {
		if spec.Type == "one2many" && spec.Relation != "" && spec.RelationField != "" {
			rcmd := NewRelatedCommand(
				odoocfg,
				spec.Relation,
				spec.RelationField,
				[]int{},
				spec.Description,
				spec.Type,
				OdooContext{},
			)
			rcmd.OriginField = fieldName
			if spec.Relation == fmt.Sprintf("%s.line", lastCommand.Model) {
				rcmd.Score = 10
			} else if strings.HasPrefix(spec.Relation, lastCommand.Model) {
				rcmd.Score = 2
			} else {
				rcmd.Score = 0
			}
			rcmds = append(rcmds, *rcmd)
		}
	}
	if len(rcmds) == 0 {
		return rcmds, fmt.Errorf("No related object found")
	}
	sort.SliceStable(rcmds, func(i, j int) bool {
		return rcmds[i].Score > rcmds[j].Score
	})
	return rcmds, nil
}

func GetRelatedCommand(model string, rcmds []RelatedCommand) (rcmd RelatedCommand, err error) {
	for _, item := range rcmds {
		if item.Model == model {
			return item, err
		}
	}
	return rcmd, fmt.Errorf("The model %s not found from the related commands", model)
}

func StringToCommand(cmd *Command, str string) error {
	str = strings.Trim(str, " ")
	parts := strLib.ToArgv(str)
	if len(parts) == 0 {
		return fmt.Errorf("the query is empty")
	}
	if cmd.Model == "" {
		cmd.Model = parts[0]
		parts = parts[1:]
	}
	fields := []string{}
	orders := []string{}
	domains := make([][]interface{}, 0)
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		domainEqualParts := strings.Split(part, "=")
		if len(domainEqualParts) == 2 {
			domainEqualValue := strings.Split(domainEqualParts[1], ",")
			domains = append(domains, []interface{}{
				domainEqualParts[0],
				"in",
				domainEqualValue,
			})
			continue
		}
		domainIlikeParts := strings.Split(part, "~")
		if len(domainIlikeParts) == 2 {
			domains = append(domains, []interface{}{
				domainIlikeParts[0],
				"ilike",
				domainIlikeParts[1],
			})
			continue
		}
		domainNotEqualParts := strings.Split(part, "!=")
		if len(domainNotEqualParts) == 2 {
			domainNotEqualValue := strings.Split(domainNotEqualParts[1], ",")
			domains = append(domains, []interface{}{
				domainNotEqualParts[0],
				"not in",
				domainNotEqualValue,
			})
			continue
		}
		domainLessThanParts := strings.Split(part, "<")
		if len(domainLessThanParts) == 2 {
			domainLessThanValue := domainLessThanParts[1]
			domains = append(domains, []interface{}{
				domainLessThanParts[0],
				"<",
				domainLessThanValue,
			})
			continue
		}
		domainLessThanOrEqualParts := strings.Split(part, "<=")
		if len(domainLessThanOrEqualParts) == 2 {
			domainLessThanOrEqualValue := domainLessThanOrEqualParts[1]
			domains = append(domains, []interface{}{
				domainLessThanOrEqualParts[0],
				"<=",
				domainLessThanOrEqualValue,
			})
			continue
		}
		domainGreatThanParts := strings.Split(part, ">")
		if len(domainGreatThanParts) == 2 {
			domainGreatThanValue := domainGreatThanParts[1]
			domains = append(domains, []interface{}{
				domainGreatThanParts[0],
				">",
				domainGreatThanValue,
			})
			continue
		}
		domainGreatThanOrEqualParts := strings.Split(part, ">=")
		if len(domainGreatThanOrEqualParts) == 2 {
			domainGreatThanOrEqualValue := domainGreatThanOrEqualParts[1]
			domains = append(domains, []interface{}{
				domainGreatThanOrEqualParts[0],
				">=",
				domainGreatThanOrEqualValue,
			})
			continue
		}
		if string(part[0]) == "+" && len(part) > 1 {
			fields = append(fields, part[1:])
			orders = append(orders, fmt.Sprintf("%v asc", part[1:]))
		} else if string(part[0]) == "-" && len(part) > 1 {
			fields = append(fields, part[1:])
			orders = append(orders, fmt.Sprintf("%v desc", part[1:]))
		} else {
			fields = append(fields, part)
		}

	}
	if len(fields) > 0 {
		cmd.Fields = fields
	}
	if len(domains) > 0 {
		cmd.Domain = domains
	}
	if len(orders) > 0 {
		cmd.Order = strings.Join(orders, ", ")
	}
	return nil
}
