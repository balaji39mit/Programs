package driver

import (
	"fmt"
	"time"

	"github.com/balaji39mit/Applications/RideSharing/vehicle"
)

type Driver struct {
	Id           int64
	UserId       int64
	VehicleId    int64
	CanOfferRide bool
	CreatedOn    time.Time
	UpdatedOn    time.Time
}

func (d Driver) String() string {
	str := fmt.Sprintf("Id : %d", d.Id)
	str += fmt.Sprintf("UserId : %d", d.UserId)
	str += fmt.Sprintf("Vehicle : %s", vehicle.Client.Display(d.VehicleId))
	return str
}
