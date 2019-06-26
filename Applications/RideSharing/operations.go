package main

import (
	"fmt"

	"github.com/balaji39mit/Applications/RideSharing/vehicle"

	"github.com/balaji39mit/Applications/RideSharing/ride"

	"github.com/balaji39mit/Applications/RideSharing/driver"
)

var OfferARide = func() {
	fmt.Println("Please find below the ride details which can be offered")
	for key, val := range driver.Client.GetAvailableDrivers() {
		fmt.Println(fmt.Sprintf("%d. %s", key+1, val))
	}
}

var StopOfferingMore = func() {
	var id int64
	fmt.Println("Enter the driver Id who don't want to offer anymore.")
	_, _ = fmt.Scanf("%d", &id)
	driver.Client.StopOffering(id)
}

var SeeAvailableRides = func() {
	var price float64
	fmt.Println("Enter price : ")
	_, _ = fmt.Scanf("%f", &price)
	rideAvailable := false
	for _, val := range driver.Client.GetAvailableDrivers() {
		if vehicle.Client.GetPrice(val.Id) <= price {
			rideAvailable = true
			fmt.Println(fmt.Sprintf("%s", val))
		}
	}

	if !rideAvailable {
		fmt.Println("Sorry no ride is available")
	}

}

var RequestRide = func() {
	fmt.Println("Enter your id, driver id and vehicle id")
	var id, did, vid int64
	_, _ = fmt.Scanf("%d %d %d", &id, &did, &vid)
	fmt.Println("Enter number of riders")
	var n int64
	_, _ = fmt.Scanf("%d", &n)
	if driver.Client.CanOffer(did, vid, n) {
		//Create Ride.
		id := ride.Client.CreateRide(id, did, vehicle.Client.GetPrice(vid))
		fmt.Println("Ride created with id : ", id)
	} else {
		fmt.Println("Sorry driver cannot offer an ride")
	}
}

var AcceptOrReject = func() {
	var id, did, vid int64
	fmt.Println("Enter id of a ride and driver id and vehicle id")
	_, _ = fmt.Scanf("%d %d %%d", &id, &did, &vid)
	if ride.Client.ValidRideForOp(id) {
		ride.Client.CreateRide(id, did, vehicle.Client.GetPrice(vid))
	} else {
		fmt.Println("Sorry invalid ride.")
	}
}

var GetRider = func() {
	fmt.Println("Enter ride id : ")
	var id int64
	fmt.Scanf("%d", &id)
	user := ride.Client.GetRider(id)
	if user != nil {
		fmt.Printf("%s", user)
	} else {
		fmt.Printf("No info")
	}
}

var GetUser = func() {
	fmt.Println("Enter user id : ")
	var id int64
	fmt.Scanf("%d", &id)
	user := ride.Client.GetUser(id)
	if user != nil {
		fmt.Printf("%s", user)
	} else {
		fmt.Printf("No info")
	}
}
