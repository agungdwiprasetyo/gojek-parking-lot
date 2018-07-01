package cmdmanager

import (
	"fmt"
	"strings"

	parking "../parking"
)

// MenuCommand is interface all registered command
type MenuCommand interface {
	ParseArgs(string) error
	Clear()
	ValidateParams() bool
	Run() (string, error)
}

// Command is object struct for manage command from user and control application
type Command struct {
	Args []string
	Menu map[string]MenuCommand
}

var (
	errInvalidParam         = fmt.Errorf("command: invalid parameter(s), please provide valid input")
	errParkingNotInitialize = fmt.Errorf("parking: please create_parking_lot first")
)

// InitCommand for register all command string for control this application
func InitCommand() *Command {
	cmd := new(Command)
	cmd.Menu = make(map[string]MenuCommand)
	cmd.Menu["create_parking_lot"] = new(CommandCreateParkingLot)
	cmd.Menu["park"] = new(CommandPark)
	cmd.Menu["leave"] = new(CommandLeave)
	cmd.Menu["status"] = new(CommandStatus)
	cmd.Menu["registration_numbers_for_cars_with_colour"] = new(CommandRegistrationNumber)
	cmd.Menu["slot_numbers_for_cars_with_colour"] = new(CommandSlotNumberCarColor)
	cmd.Menu["slot_number_for_registration_number"] = new(CommandSlotNumberCarNumber)
	return cmd
}

// Run command with parameter string entered by user
func (this *Command) Run(command string) string {
	cmds := strings.SplitN(command, " ", 2)
	menu := cmds[0]
	cmdChild, ok := this.Menu[menu]
	if cmdChild == nil || !ok {
		return fmt.Sprintf("\x1b[31;1m%s: command not found\x1b[0m", menu)
	}

	// clearing command argument(s)
	cmdChild.Clear()

	// capture command argument(s)
	if len(cmds) > 1 {
		cmdChild.ParseArgs(cmds[1])
	}

	valid := cmdChild.ValidateParams()
	if !valid {
		return fmt.Sprintf("\x1b[31;1m%s: invalid parameter(s), please provide valid input\x1b[0m", menu)
	}

	park := parking.Get()
	if park == nil && menu != "create_parking_lot" {
		return fmt.Sprintf("\x1b[31;1m%v\x1b[0m", errParkingNotInitialize)
	}

	output, err := cmdChild.Run()
	if err != nil {
		return fmt.Sprintf("\x1b[31;1m%v\x1b[0m", err)
	}
	return output
}
