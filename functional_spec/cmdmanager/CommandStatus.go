package cmdmanager

import (
	"fmt"
	"strings"

	parking "../parking"
)

type CommandStatus struct {
	Command
}

func (this *CommandStatus) ParseArgs(args string) error {
	return nil
}

func (this *CommandStatus) Clear() {
	this.Args = nil
}

func (this *CommandStatus) ValidateParams() bool {
	return true
}

func (this *CommandStatus) Run() (string, error) {
	var list = []string{fmt.Sprintf("%-12s%-20s%-10s", "Slot No.", "Registration No", "Color")}
	filledSlots := parking.Get().GetFilledSlots()
	for _, sl := range filledSlots {
		cr := sl.Car
		list = append(list, fmt.Sprintf("%-12v%-20v%-10v", sl.Index, cr.Number, cr.Color))
	}
	output := strings.Join(list, "\n")
	return output, nil
}
