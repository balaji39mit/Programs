package ride

import "github.com/balaji39mit/Applications/RideSharing/user"

type IRide interface {
	RequestRide(int64)
	CreateRide(int64, int64, float64) int64
	ValidRideForOp(int64) bool
	GetRider(int64) *user.User
	GetUser(id int64) *user.User
}
