package parking

import (
	"fmt"
	"strings"

	car "../car"
	slot "../slot"
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
	parking.Save()
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

func (this *Parking) AddCar(cr car.Car) (*slot.Slot, error) {
	sl, err := this.FindNearestSlot()
	if err != nil {
		return nil, err
	}
	if err := sl.Allocate(cr); err != nil {
		return nil, err
	}
	return sl, nil
}

func (this *Parking) RemoveCar(cr car.Car) {
	for i, sl := range this.Slots {
		if !sl.IsFree() && sl.Car.IsEqual(cr) {
			this.Slots[i].Free()
		}
	}
}

func (this *Parking) GetFilledSlots() (filledSlots []*slot.Slot) {
	for _, sl := range this.Slots {
		if !sl.IsFree() {
			filledSlots = append(filledSlots, sl)
		}
	}
	return
}

func (this *Parking) GetSlotsByCarColor(carColor string) (slots []*slot.Slot) {
	for _, sl := range this.Slots {
		if !sl.IsFree() {
			if strings.ToLower(sl.GetCarColor()) == strings.ToLower(carColor) {
				slots = append(slots, sl)
			}
		}
	}
	return
}

func (this *Parking) RemoveCarBySlot(slotNumber uint) error {
	for i, sl := range this.Slots {
		if sl.Index == slotNumber {
			this.Slots[i].Car = nil
			return nil
		}
	}
	return fmt.Errorf("Slot %d not found", slotNumber)
}

func (this *Parking) RemoveSlot() {

}
