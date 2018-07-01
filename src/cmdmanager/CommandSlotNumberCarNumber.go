package cmdmanager

import (
	"fmt"
	"strings"

	parking "../parking"
)

type CommandSlotNumberCarNumber struct {
	Command
	CarNumber string
}

func (this *CommandSlotNumberCarNumber) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	this.CarNumber = this.Args[0]
	return nil
}

func (this *CommandSlotNumberCarNumber) Clear() {
	this.Args = nil
	this.CarNumber = ""
}

func (this *CommandSlotNumberCarNumber) ValidateParams() bool {
	return len(this.Args) == 1 && this.Args[0] != ""
}

func (this *CommandSlotNumberCarNumber) Run() (string, error) {
	var output string
	slot := parking.Get().GetSlotByCarNumber(this.CarNumber)
	if slot == nil {
		return "Not found", nil
	}
	output = fmt.Sprint(slot.Index)
	return output, nil
}
