package parking_lot

import "fmt"

type Car struct {
	RegistrationNumber string
	Color              string
}

func (c Car) String() string {
	return fmt.Sprintf("RegNo : %s, Color : %s", c.RegistrationNumber, c.Color)
}
