package slot

import (
	"fmt"

	vehicle "../vehicle"
)

type Slot struct {
	Index   uint
	Vehicle *vehicle.Vehicle
}

func New() *Slot {
	return &Slot{}
}

func (this *Slot) Allocate(vh vehicle.Vehicle) error {
	if this.Vehicle != nil {
		return fmt.Errorf("slot: Slot already allocated")
	}
	this.Vehicle = &vh
	return nil
}

func (this *Slot) Free() {
	this.Vehicle = nil
}

func (this *Slot) IsFree() bool {
	return this.Vehicle == nil
}
