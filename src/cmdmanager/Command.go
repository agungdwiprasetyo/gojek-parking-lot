package cmdmanager

import (
	"fmt"
	"strings"
)

type MenuCommand interface {
	ParseArgs(string) error
	ValidateParams() bool
	Run() (string, error)
}

type Command struct {
	Args    []string
	Manager map[string]MenuCommand
}

var (
	errInvalidParam = fmt.Errorf("command: invalid parameter(s), please provide valid input")
)

func InitCommand() *Command {
	cmd := new(Command)
	cmd.Manager = make(map[string]MenuCommand)
	cmd.Manager["create_parking_lot"] = new(CommandCreateParkingLot)
	return cmd
}

func (this *Command) Clear() {
	this.Args = []string{}
}

func (this *Command) Run(command string) string {
	cmds := strings.SplitN(command, " ", 2)
	menu := cmds[0]
	cmdChild, ok := this.Manager[menu]
	if cmdChild == nil || !ok {
		return fmt.Sprintf("\x1b[31;1m%s: command not found\x1b[0m", menu)
	}
	this.Clear()

	// capture command argument(s)
	if len(cmds) > 1 {
		cmdChild.ParseArgs(cmds[1])
	}

	valid := cmdChild.ValidateParams()
	if !valid {
		return fmt.Sprintf("\x1b[31;1m%s: invalid parameter(s), please provide valid input\x1b[0m", menu)
	}

	output, err := cmdChild.Run()
	if err != nil {
		return fmt.Sprintf("\x1b[31;1m%v\x1b[0m", err)
	}
	return output
}
