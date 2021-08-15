package data

import (
	"fmt"

	"github.com/chermed/kodoo/pkg/odoo"
)

type CommandsHistory struct {
	Commands []*odoo.Command
	Index    int
}

func (commandHistory *CommandsHistory) AddCommand(cmd *odoo.Command) {
	if commandHistory.GetCommandsSize() > 0 && commandHistory.Index < commandHistory.GetCommandsSize()-1 {
		commandHistory.Commands = commandHistory.Commands[:commandHistory.Index+1]
	}
	commandHistory.Commands = append(commandHistory.Commands, cmd)
	commandHistory.Index = len(commandHistory.Commands) - 1
}
func (commandHistory *CommandsHistory) HasCommand() bool {
	if len(commandHistory.Commands) > 0 {
		return true
	}
	return false
}
func (commandHistory *CommandsHistory) GetCommandsSize() int {
	return len(commandHistory.Commands)
}
func (commandHistory *CommandsHistory) GetCommand() (*odoo.Command, error) {
	if !commandHistory.HasCommand() {
		return &odoo.Command{}, fmt.Errorf("Please run a command before this action")
	}
	return commandHistory.Commands[commandHistory.Index], nil
}
func (commandHistory *CommandsHistory) GetCommandCopy() (odoo.Command, error) {
	if len(commandHistory.Commands) == 0 {
		return odoo.Command{}, fmt.Errorf("Please run a command before this action")
	}
	cmd := *commandHistory.Commands[commandHistory.Index]
	cmd.FieldsUpdated = false
	return cmd, nil
}
func (commandHistory *CommandsHistory) GoToNextCommand() error {
	return commandHistory.goToIndex(commandHistory.Index + 1)
}
func (commandHistory *CommandsHistory) GoToPreviousCommand() error {
	return commandHistory.goToIndex(commandHistory.Index - 1)
}
func (commandHistory *CommandsHistory) goToIndex(index int) error {
	if commandHistory.GetCommandsSize() == 0 {
		return fmt.Errorf("No command found in the commandHistory")
	}
	if index >= len(commandHistory.Commands) {
		commandHistory.Index = len(commandHistory.Commands) - 1
		return fmt.Errorf("You are on the last command")
	} else if index < 0 {
		commandHistory.Index = 0
		return fmt.Errorf("You are on the first command")
	} else {
		commandHistory.Index = index
		return nil
	}
}
