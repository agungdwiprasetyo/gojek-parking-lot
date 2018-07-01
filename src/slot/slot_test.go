package slot

import (
	"testing"

	car "../car"
)

func TestSlot_Allocate(t *testing.T) {
	type fields struct {
		Index uint
		Car   *car.Car
	}
	type args struct {
		cr car.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"TestCase 1: Slot already allocated",
			fields{Index: 2, Car: &car.Car{Number: "BE4508GE", Color: "Red"}},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			true,
		},
		{
			"TestCase 2: Slot free",
			fields{Index: 2, Car: nil},
			args{cr: car.Car{Number: "BE4508GE", Color: "Red"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Slot{
				Index: tt.fields.Index,
				Car:   tt.fields.Car,
			}
			if err := this.Allocate(tt.args.cr); (err != nil) != tt.wantErr {
				t.Errorf("Slot.Allocate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlot_Free(t *testing.T) {
	type fields struct {
		Index uint
		Car   *car.Car
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"TestCase 1: Slot is free",
			fields{Index: 1, Car: &car.Car{Number: "BE4508GE", Color: "Red"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Slot{
				Index: tt.fields.Index,
				Car:   tt.fields.Car,
			}
			this.Free()
			if this.Car != nil {
				t.Errorf("Slot.Free() car is not nil, want nil")
			}
		})
	}
}

func TestSlot_IsFree(t *testing.T) {
	type fields struct {
		Index uint
		Car   *car.Car
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Slot{
				Index: tt.fields.Index,
				Car:   tt.fields.Car,
			}
			if got := this.IsFree(); got != tt.want {
				t.Errorf("Slot.IsFree() = %v, want %v", got, tt.want)
			}
		})
	}
}
