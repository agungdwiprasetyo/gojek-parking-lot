package parking

import (
	"fmt"

	slot "../slot"
	vehicle "../vehicle"
)

type Parking struct {
	Capacity uint
	Slots    []*slot.Slot
}

var (
	saved      *Parking
	startIndex = 1
)

func New(capacity uint) *Parking {
	parking := new(Parking)
	parking.Capacity = capacity
	parking.Slots = make([]*slot.Slot, capacity)
	idx := startIndex
	for i := range parking.Slots {
		parking.Slots[i] = &slot.Slot{Index: uint(idx)}
		idx++
	}
	saved = parking
	return parking
}

// Get for get saved data
func Get() *Parking {
	return saved
}

func (this *Parking) Save() {
	saved = this
}

func (this *Parking) FindNearestSlot() (*slot.Slot, error) {
	for _, sl := range this.Slots {
		if sl.IsFree() {
			return sl, nil
		}
	}
	return nil, fmt.Errorf("Sorry, parking lot is full")
}

func (this *Parking) AddVehicle(vh vehicle.Vehicle) (*slot.Slot, error) {
	sl, err := this.FindNearestSlot()
	if err != nil {
		return nil, err
	}
	if err := sl.Allocate(vh); err != nil {
		return nil, err
	}
	return sl, nil
}

func (this *Parking) RemoveVehicle(vh vehicle.Vehicle) {
	for i, sl := range this.Slots {
		if sl.Vehicle != nil && sl.Vehicle.IsEqual(vh) {
			this.Slots[i].Free()
		}
	}
}

func (this *Parking) GetFilledSlots() []*slot.Slot {
	var filledSlots []*slot.Slot
	for _, sl := range this.Slots {
		if !sl.IsFree() {
			filledSlots = append(filledSlots, sl)
		}
	}
	return filledSlots
}

func (this *Parking) RemoveVehicleBySlot(slotNumber uint) {
	this.Slots[slotNumber-1].Vehicle = nil
}

func (this *Parking) RemoveSlot() {

}
