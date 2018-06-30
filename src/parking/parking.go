package parking

import (
	"fmt"

	slot "../slot"
	vehicle "../vehicle"
)

type Parking struct {
	Capacity        uint
	AllocationIndex uint
	Slots           []*slot.Slot
}

var (
	startIndex = 1
)

func New(capacity uint) *Parking {
	parking := new(Parking)
	parking.Capacity = capacity
	parking.AllocationIndex = uint(startIndex)
	parking.Slots = make([]*slot.Slot, capacity)
	idx := startIndex
	for i := range parking.Slots {
		parking.Slots[i] = &slot.Slot{Index: uint(idx)}
		idx++
	}
	return parking
}

func (this *Parking) FindNearestSlot() (*slot.Slot, error) {
	for _, sl := range this.Slots {
		if sl.IsFree() {
			return sl, nil
		}
	}
	return nil, fmt.Errorf("Sorry, parking lot is full")
}

func (this *Parking) AddVehicle(vh vehicle.Vehicle) error {
	sl, err := this.FindNearestSlot()
	if err != nil {
		return err
	}
	err = sl.Allocate(vh)
	return err
}

func (this *Parking) RemoveVehicle(vh vehicle.Vehicle) {
	for i, sl := range this.Slots {
		if sl.Vehicle != nil && sl.Vehicle.IsEqual(vh) {
			this.Slots[i].Vehicle = nil
		}
	}
}

func (this *Parking) RemoveVehicleBySlot(slotNumber uint) {}

func (this *Parking) RemoveSlot() {

}
