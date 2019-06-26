package driver

import (
	"sync"
	"time"

	"github.com/balaji39mit/Applications/RideSharing/vehicle"
)

var Client IDriver

type Impl struct{}

var drivers []*Driver

func init() {
	Client = Impl{}
	drivers = []*Driver{
		{
			Id:           1,
			UserId:       1,
			VehicleId:    1,
			CanOfferRide: true,
			CreatedOn:    time.Now(),
			UpdatedOn:    time.Now(),
		},
		{
			Id:           2,
			UserId:       2,
			VehicleId:    2,
			CanOfferRide: true,
			CreatedOn:    time.Now(),
			UpdatedOn:    time.Now(),
		},
	}
}

func (d Impl) GetAvailableDrivers() []Driver {
	res := make([]Driver, 0)
	for _, val := range drivers {
		if val.CanOfferRide {
			res = append(res, *val)
		}
	}
	return res
}

func (d Impl) StopOffering(id int64) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	for _, val := range drivers {
		if val.Id == id {
			val.CanOfferRide = false
			return
		}
	}
}

func (d Impl) CanOffer(driverId, vehicleId int64, n int64) bool {
	for _, val := range drivers {
		if val.Id == driverId {
			if val.CanOfferRide {
				//Treating n as 1 for now.
				return vehicle.Client.CanServeRide(vehicleId, n)
			}
		}
	}
	return false
}
