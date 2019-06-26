package ride

import (
	"sync"
	"time"

	"github.com/balaji39mit/Applications/RideSharing/user"
)

var (
	Client IRide
	rides  []*Rides
)

type Impl struct {
}

func init() {
	Client = Impl{}
	rides = make([]*Rides, 0)
}

func (r Impl) RequestRide(n int64) {

}

func (r Impl) CreateRide(rid int64, did int64, price float64) int64 {
	ride := &Rides{
		Id:        int64(len(rides)) + 1,
		RiderId:   rid,
		DriverId:  did,
		Price:     price,
		Status:    Initiated,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	rides = append(rides, ride)
	return ride.Id
}

func (r Impl) getRide(id int64) *Rides {
	for _, val := range rides {
		if val.Id == id {
			return val
		}
	}
	return nil
}

func (r Impl) ValidRideForOp(id int64) bool {
	for _, val := range rides {
		if val.Id == id {
			return val.Status == Initiated
		}
	}
	return false
}

func (r Impl) GetRider(id int64) *user.User {
	if r.ValidRideForOp(id) {
		ride := r.getRide(id)
		if ride != nil {
			return user.Client.Get(ride.RiderId)
		}
	}
	return nil
}

func (r Impl) GetUser(id int64) *user.User {
	if r.ValidRideForOp(id) {
		ride := r.getRide(id)
		if ride != nil {
			return user.Client.Get(ride.DriverId)
		}
	}
	return nil
}
