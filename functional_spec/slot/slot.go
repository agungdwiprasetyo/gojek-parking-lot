package slot

import (
	"fmt"

	car "../car"
)

// Slot is a cleared area that is intended for parking car, with identity serial index number.
type Slot struct {
	Index uint
	Car   *car.Car
}

func New() *Slot {
	return &Slot{}
}

func (this *Slot) Allocate(cr car.Car) error {
	if this.Car != nil {
		return fmt.Errorf("slot: Slot already allocated")
	}
	this.Car = &cr
	return nil
}

func (this *Slot) GetCarNumber() string {
	return this.Car.Number
}

func (this *Slot) GetCarColor() string {
	return this.Car.Color
}

func (this *Slot) Free() {
	this.Car = nil
}

func (this *Slot) IsFree() bool {
	return this.Car == nil
}
