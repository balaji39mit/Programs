package ride

import "time"

type RideStatus uint8

const (
	Initiated  RideStatus = 1
	Accepted   RideStatus = 2
	Rejected   RideStatus = 3
	InProgress RideStatus = 4
	Completed  RideStatus = 5
)

type Rides struct {
	Id        int64
	RiderId   int64
	DriverId  int64
	Price     float64
	Status    RideStatus
	CreatedOn time.Time
	UpdatedOn time.Time
}
