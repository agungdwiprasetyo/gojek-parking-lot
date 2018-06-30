package parking

import (
	"reflect"
	"testing"

	slot "../slot"
	vehicle "../vehicle"
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
				Capacity:        3,
				AllocationIndex: 1,
				Slots: []*slot.Slot{
					{
						Index:   1,
						Vehicle: nil,
					},
					{
						Index:   2,
						Vehicle: nil,
					},
					{
						Index:   3,
						Vehicle: nil,
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

func TestParking_AddVehicle(t *testing.T) {
	type fields struct {
		Capacity        uint
		AllocationIndex uint
		Slots           []*slot.Slot
	}
	type args struct {
		vh vehicle.Vehicle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"TestCase 1: Parking slot is not full",
			fields{
				Capacity:        3,
				AllocationIndex: 1,
				Slots: []*slot.Slot{
					{
						Index:   1,
						Vehicle: nil,
					},
					{
						Index:   2,
						Vehicle: nil,
					},
					{
						Index:   3,
						Vehicle: nil,
					},
				},
			},
			args{vh: vehicle.Vehicle{Number: "BE4508GE", Color: "Red"}},
			false,
		},
		{
			"TestCase 2: Parking slots is full",
			fields{
				Capacity:        3,
				AllocationIndex: 1,
				Slots: []*slot.Slot{
					{
						Index:   1,
						Vehicle: &vehicle.Vehicle{Number: "BE1000GE", Color: "Red"},
					},
					{
						Index:   2,
						Vehicle: &vehicle.Vehicle{Number: "BE2000GE", Color: "Red"},
					},
					{
						Index:   3,
						Vehicle: &vehicle.Vehicle{Number: "BE3000GE", Color: "Red"},
					},
				},
			},
			args{vh: vehicle.Vehicle{Number: "BE4508GE", Color: "Red"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Parking{
				Capacity:        tt.fields.Capacity,
				AllocationIndex: tt.fields.AllocationIndex,
				Slots:           tt.fields.Slots,
			}
			if err := this.AddVehicle(tt.args.vh); (err != nil) != tt.wantErr {
				t.Errorf("Parking.AddVehicle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
