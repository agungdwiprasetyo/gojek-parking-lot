package cmdmanager

import (
	"strings"

	parking "../parking"
)

type CommandRegistrationNumber struct {
	Command
	CarColor string
}

func (this *CommandRegistrationNumber) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	this.CarColor = this.Args[0]
	return nil
}

func (this *CommandRegistrationNumber) Clear() {
	this.Args = nil
	this.CarColor = ""
}

func (this *CommandRegistrationNumber) ValidateParams() bool {
	return len(this.Args) == 1 && this.Args[0] != ""
}

func (this *CommandRegistrationNumber) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(this.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, sl.GetCarNumber())
	}
	output = strings.Join(list, ", ")
	return output, nil
}
