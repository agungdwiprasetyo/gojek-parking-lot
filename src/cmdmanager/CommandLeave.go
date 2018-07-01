package cmdmanager

import (
	"fmt"
	"strconv"
	"strings"

	parking "../parking"
)

type CommandLeave struct {
	Command
	SlotNumber uint
}

func (this *CommandLeave) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	value, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		return errInvalidParam
	}
	this.SlotNumber = uint(value)
	return nil
}

func (this *CommandLeave) Clear() {
	this.Args = nil
	this.SlotNumber = 0
}

func (this *CommandLeave) ValidateParams() bool {
	return len(this.Args) == 1
}

func (this *CommandLeave) Run() (string, error) {
	var output string
	park := parking.Get()
	slotNumber := this.SlotNumber
	if err := park.RemoveCarBySlot(slotNumber); err != nil {
		return output, err
	}
	output = fmt.Sprintf("Slot number %d is free", slotNumber)
	return output, nil
}
