package slot

import (
	vehicle "../vehicle"
)

type Slot struct {
	Index   uint
	Vehicle *vehicle.Vehicle
}

func New(index int, vehicle vehicle.Vehicle) *Slot {
	return &Slot{Index: uint(index), Vehicle: &vehicle}
}
