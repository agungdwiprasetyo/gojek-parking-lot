package car

import "strings"

// Car is a basic object. Car will fill a slot and reduce parking lot slot capacity.
type Car struct {
	Number string
	Color  string
}

// New is package Car constructor
func New(number, color string) *Car {
	return &Car{
		Number: number,
		Color:  color,
	}
}

// IsEqual is method for check deeply equal between 2 Car object
func (this *Car) IsEqual(cr Car) bool {
	return (this.Number == cr.Number) &&
		(strings.ToLower(this.Color) == strings.ToLower(cr.Color))
}
