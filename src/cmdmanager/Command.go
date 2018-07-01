package cmdmanager

import (
	"fmt"
	"strings"

	parking "../parking"
)

type MenuCommand interface {
	ParseArgs(string) error
	Clear()
	ValidateParams() bool
	Run() (string, error)
}

type Command struct {
	Args []string
	Menu map[string]MenuCommand
}

var (
	errInvalidParam         = fmt.Errorf("command: invalid parameter(s), please provide valid input")
	errParkingNotInitialize = fmt.Errorf("parking: please create_parking_lot first")
)

func InitCommand() *Command {
	cmd := new(Command)
	cmd.Menu = make(map[string]MenuCommand)
	cmd.Menu["create_parking_lot"] = new(CommandCreateParkingLot)
	cmd.Menu["park"] = new(CommandPark)
	cmd.Menu["leave"] = new(CommandLeave)
	cmd.Menu["status"] = new(CommandStatus)
	cmd.Menu["registration_numbers_for_cars_with_colour"] = new(CommandRegistrationNumber)
	return cmd
}

func (this *Command) Run(command string) string {
	cmds := strings.SplitN(command, " ", 2)
	menu := cmds[0]
	cmdChild, ok := this.Menu[menu]
	if cmdChild == nil || !ok {
		return fmt.Sprintf("\x1b[31;1m%s: command not found\x1b[0m", menu)
	}

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
	if park == nil {
		return fmt.Sprintf("\x1b[31;1m%v\x1b[0m", errParkingNotInitialize)
	}

	output, err := cmdChild.Run()
	if err != nil {
		return fmt.Sprintf("\x1b[31;1m%v\x1b[0m", err)
	}
	return output
}
