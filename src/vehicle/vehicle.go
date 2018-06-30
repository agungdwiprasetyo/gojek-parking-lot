package vehicle

// Vehicle is object for
type Vehicle struct {
	Number string
	Color  string
}

// New is package vehicle constructor
func New(number, color string) *Vehicle {
	return &Vehicle{
		Number: number,
		Color:  color,
	}
}

// DeepEqual is method for check deeply equal between 2 vehicle object
func (this *Vehicle) DeepEqual(vehicle Vehicle) bool {
	return (this.Number == vehicle.Number) && (this.Color == vehicle.Color)
}
