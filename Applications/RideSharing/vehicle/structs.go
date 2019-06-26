package vehicle

import "time"

type VehicleType uint8

const (
	TwoWheeler   VehicleType = 2
	ThreeWheeler VehicleType = 3
	FourWheeler  VehicleType = 4
)

type Vehicles struct {
	Id         int64
	Name       string
	Type       VehicleType
	Capacity   int64
	BasePrice  float64
	PricePerKM float64
	CreatedOn  time.Time
	UpdatedOn  time.Time
}
