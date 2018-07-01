package cmdmanager

import "testing"

func TestCommand_Run(t *testing.T) {
	type fields struct {
		Args []string
		Menu map[string]MenuCommand
	}
	type args struct {
		command string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"TestCase 1: Test create parking lot with capacity 6 (create_parking_lot 6)",
			fields{
				Menu: map[string]MenuCommand{"create_parking_lot": &CommandCreateParkingLot{}},
			},
			args{command: "create_parking_lot 6"},
			"Created a parking lot with 6 slots",
		},
		{
			"TestCase 2: Test undefined command (e.g foo)",
			fields{
				Menu: map[string]MenuCommand{"foo": nil},
			},
			args{command: "foo"},
			"\x1b[31;1mfoo: command not found\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Command{
				Args: tt.fields.Args,
				Menu: tt.fields.Menu,
			}
			if got := this.Run(tt.args.command); got != tt.want {
				t.Errorf("Command.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
