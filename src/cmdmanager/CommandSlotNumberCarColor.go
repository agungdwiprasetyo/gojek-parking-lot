package cmdmanager

import (
	"fmt"
	"strings"

	parking "../parking"
)

type CommandSlotNumberCarColor struct {
	Command
	CarColor string
}

func (this *CommandSlotNumberCarColor) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	this.CarColor = this.Args[0]
	return nil
}

func (this *CommandSlotNumberCarColor) Clear() {
	this.Args = nil
	this.CarColor = ""
}

func (this *CommandSlotNumberCarColor) ValidateParams() bool {
	return len(this.Args) == 1 && this.Args[0] != ""
}

func (this *CommandSlotNumberCarColor) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(this.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, fmt.Sprint(sl.Index))
	}
	output = strings.Join(list, ", ")
	return output, nil
}
