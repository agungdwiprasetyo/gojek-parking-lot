package cmdmanager

import (
	"fmt"
	"strconv"
	"strings"

	parking "../parking"
)

type CommandCreateParkingLot struct {
	Command
	Capacity uint
}

func (this *CommandCreateParkingLot) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	value, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		return errInvalidParam
	}
	this.Capacity = uint(value)
	return nil
}

func (this *CommandCreateParkingLot) Clear() {
	this.Args = nil
	this.Capacity = 0
}

func (this *CommandCreateParkingLot) ValidateParams() bool {
	return len(this.Args) == 1
}

func (this *CommandCreateParkingLot) Run() (string, error) {
	var output string
	obj := parking.New(this.Capacity)
	if obj == nil {
		return output, fmt.Errorf("Error")
	}
	output = fmt.Sprintf("Created a parking lot with %d slots", this.Capacity)
	return output, nil
}
