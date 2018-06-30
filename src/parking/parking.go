package parking

import (
	"fmt"

	"../slot"
)

type Parking struct {
	Capacity uint
	Slots    []slot.Slot
}

func New(capacity uint) *Parking {
	fmt.Printf("Created a parking lot with %d slots", capacity)
	return &Parking{Capacity: capacity}
}

func (this *Parking) AddSlot(slot slot.Slot) error {
	if len(this.Slots) > int(this.Capacity) {
		return fmt.Errorf("Sorry, parking lot is full")
	}
	this.Slots = append(this.Slots, slot)
	return nil
}

func (this *Parking) RemoveSlot() {

}
