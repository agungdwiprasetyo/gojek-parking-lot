package cmdmanager

import (
	"testing"

	car "../car"
)

func TestCommandPark_ValidateParams(t *testing.T) {
	type fields struct {
		Command Command
		Car     *car.Car
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"TestCase 1: Only two arguments (actualy valid)",
			fields{Command: Command{Args: []string{"BE4508GE", "Red"}}},
			true,
		},
		{
			"TestCase 2: Total arguments not same as two (actualy invalid)",
			fields{Command: Command{Args: []string{"BE4508GE", "Red", "blue"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CommandPark{
				Command: tt.fields.Command,
				Car:     tt.fields.Car,
			}
			if got := this.ValidateParams(); got != tt.want {
				t.Errorf("CommandPark.ValidateParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
