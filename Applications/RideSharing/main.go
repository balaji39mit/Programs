package main

import "fmt"

var Features map[string]func()

func init() {
	Features = make(map[string]func(), 0)
	Features = map[string]func(){
		"OfferARide":        OfferARide,
		"StopOfferingMore":  StopOfferingMore,
		"SeeAvailableRides": SeeAvailableRides,
		"RequestRide":       RequestRide,
		"AcceptOrReject":    AcceptOrReject,
		"GetRider":          GetRider,
		"GetUser":           GetUser,
	}
}

func main() {
	fmt.Println("Ride Sharing Application..")
	for {
		fmt.Println("Please choose one of the operation : ")
		count := 1
		for key := range Features {
			fmt.Println(count, ".", key)
			count++
		}
		var operation string
		_, _ = fmt.Scanf("%s", &operation)
		if val, ok := Features[operation]; ok {
			val()
		} else {
			fmt.Println("Invalid operation.. Exiting the application.")
			break
		}
		fmt.Println("Operation completed.")
	}
}
