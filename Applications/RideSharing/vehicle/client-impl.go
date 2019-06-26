package vehicle

import (
	"time"

	"fmt"
)

var Client IVehicle

type Impl struct {
}

var vehicles []*Vehicles

func init() {
	Client = Impl{}
	vehicles = []*Vehicles{
		{
			Id:         1,
			Name:       "Honda",
			Type:       FourWheeler,
			Capacity:   5,
			BasePrice:  100,
			PricePerKM: 10,
			CreatedOn:  time.Now(),
			UpdatedOn:  time.Now(),
		},
		{
			Id:         2,
			Name:       "Enfield",
			Type:       TwoWheeler,
			Capacity:   1,
			BasePrice:  10,
			PricePerKM: 2,
			CreatedOn:  time.Now(),
			UpdatedOn:  time.Now(),
		},
	}
}

func (v Impl) CanServeRide(id int64, n int64) bool {
	for _, val := range vehicles {
		if val.Id == id {
			return n <= val.Capacity
		}
	}
	return false
}

func (v Impl) getVehicle(id int64) *Vehicles {
	for _, val := range vehicles {
		if val.Id == id {
			return val
		}
	}
	return nil
}

//Assuming ride is offered for per km.
func (v Impl) GetPrice(id int64) float64 {
	for _, val := range vehicles {
		if val.Id == id {
			return val.BasePrice + val.PricePerKM
		}
	}
	return 0
}

func (v Impl) Display(id int64) string {
	vehicle := v.getVehicle(id)
	if vehicle != nil {
		str := fmt.Sprintf("Id : %d", vehicle.Id)
		str += fmt.Sprintf("Price : %f", vehicle.BasePrice+vehicle.PricePerKM)
		return str
	}
	return ""
}
