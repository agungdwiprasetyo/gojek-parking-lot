package parking

import (
	"reflect"
	"testing"

	car "../car"
	slot "../slot"
)

func TestNew(t *testing.T) {
	type args struct {
		capacity uint
	}

	tests := []struct {
		name string
		args args
		want *Parking
	}{
		{
			"TestCase 1: ",
			args{capacity: 3},
			&Parking{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Index: 1,
						Car:   nil,
					},
					{
						Index: 2,
						Car:   nil,
					},
					{
						Index: 3,
						Car:   nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestParking_FindNearestSlot(t *testing.T) {
	type fields struct {
		Capacity        uint
		AllocationIndex uint
		Slots           []*slot.Slot
	}
	tests := []struct {
		name    string
		fields  fields
		want    *slot.Slot
		wantErr bool
	}{
		{
			"TestCase 1",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Index: 1,
						Car:   nil,
					},
					{
						Index: 2,
						Car:   &car.Car{Number: "BE4508GE", Color: "Red"},
					},
					{
						Index: 3,
						Car:   nil,
					},
				},
			},
			&slot.Slot{
				Index: 1,
				Car:   nil,
			},
			false,
		},
		{
			"TestCase 2",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Index: 1,
						Car:   &car.Car{Number: "BE4508GE", Color: "Red"},
					},
					{
						Index: 2,
						Car:   nil,
					},
					{
						Index: 3,
						Car:   nil,
					},
				},
			},
			&slot.Slot{
				Index: 2,
				Car:   nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			got, err := this.FindNearestSlot()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parking.FindNearestSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parking.FindNearestSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParking_AddCar(t *testing.T) {
	type fields struct {
		Capacity uint
		Slots    []*slot.Slot
	}
	type args struct {
		cr car.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *slot.Slot
		wantErr bool
	}{
		{
			"TestCase 1: Parking slots is not full",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Index: 1,
						Car:   nil,
					},
					{
						Index: 2,
						Car:   nil,
					},
					{
						Index: 3,
						Car:   nil,
					},
				},
			},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			&slot.Slot{
				Index: 1,
				Car:   &car.Car{Number: "BE4508GE", Color: "Red"},
			},
			false,
		},
		{
			"TestCase 2: Parking slots is full",
			fields{
				Capacity: 3,
				Slots: []*slot.Slot{
					{
						Index: 1,
						Car:   &car.Car{Number: "BE1000GE", Color: "Red"},
					},
					{
						Index: 2,
						Car:   &car.Car{Number: "BE2000GE", Color: "Red"},
					},
					{
						Index: 3,
						Car:   &car.Car{Number: "BE3000GE", Color: "Red"},
					},
				},
			},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity: tt.fields.Capacity,
				Slots:    tt.fields.Slots,
			}
			got, err := this.AddCar(tt.args.cr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parking.AddCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parking.AddCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
