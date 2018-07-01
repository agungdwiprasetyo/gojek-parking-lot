package car

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		number string
		color  string
	}
	tests := []struct {
		name string
		args args
		want *Car
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.number, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_IsEqual(t *testing.T) {
	type fields struct {
		Number string
		Color  string
	}
	type args struct {
		cr Car
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Testcase 1: Test equal color. Test same color.",
			fields{
				Number: "BE4508GE",
				Color:  "Red",
			},
			args{
				cr: Car{Number: "BE4508GE", Color: "Red"},
			},
			true,
		},
		{
			"Testcase 2: Test equal color. Test different color.",
			fields{
				Number: "BE4508GE",
				Color:  "Red",
			},
			args{
				cr: Car{Number: "BE4508GE", Color: "Blue"},
			},
			false,
		},
		{
			"Testcase 3: Test equal color. Test different color.",
			fields{
				Number: "BE4508GE",
				Color:  "Red",
			},
			args{
				cr: Car{Number: "BE4508GE", Color: "red"},
			},
			true,
		},
		{
			"Testcase 4: Test equal number. Test different number.",
			fields{
				Number: "BE4508GE",
				Color:  "Red",
			},
			args{
				cr: Car{Number: "be4508ge", Color: "red"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Car{
				Number: tt.fields.Number,
				Color:  tt.fields.Color,
			}
			if got := this.IsEqual(tt.args.cr); got != tt.want {
				t.Errorf("Car.IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
