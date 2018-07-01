package slot

import (
	"testing"

	vehicle "../vehicle"
)

func TestSlot_Allocate(t *testing.T) {
	type fields struct {
		Index   uint
		Vehicle *vehicle.Vehicle
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Slot{
				Index:   tt.fields.Index,
				Vehicle: tt.fields.Vehicle,
			}
			if err := this.Allocate(tt.args.vh); (err != nil) != tt.wantErr {
				t.Errorf("Slot.Allocate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
