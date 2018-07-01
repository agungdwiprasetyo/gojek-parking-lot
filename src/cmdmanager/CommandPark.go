package cmdmanager

import (
	"fmt"
	"strings"

	car "../car"
	parking "../parking"
)

type CommandPark struct {
	Command
	Car *car.Car
}

func (this *CommandPark) ParseArgs(args string) error {
	this.Args = strings.Split(args, " ")
	if !this.ValidateParams() {
		return errInvalidParam
	}
	this.Car = car.New(this.Args[0], this.Args[1])
	return nil
}

func (this *CommandPark) Clear() {
	this.Args = nil
	this.Car = nil
}

func (this *CommandPark) ValidateParams() bool {
	return len(this.Args) == 2
}

func (this *CommandPark) Run() (string, error) {
	var output string
	sl, err := parking.Get().AddCar(*this.Car)
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("Allocated slot number: %d", sl.Index)
	return output, err
}
