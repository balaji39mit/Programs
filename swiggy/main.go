package main

import (
	"fmt"

	pl "github.com/balaji39mit/swiggy/parking_lot"
)

func main() {
	var op string
	for run := true; run; {
		fmt.Println("\nEnter any operation")
		_, _ = fmt.Scanf("%s", &op)
		switch op {
		case "create":
			var n int
			_, _ = fmt.Scanf("%d", &n)
			pl.ParkingLotService.Create(n)
		case "park":
			var number, color string
			_, _ = fmt.Scanf("%s %s", &number, &color)
			pl.ParkingLotService.Park(pl.Car{RegistrationNumber: number, Color: color})
		case "leave":
			var n int
			_, _ = fmt.Scanf("%d", &n)
			pl.ParkingLotService.Vacate(n)
		case "status_allocated":
			pl.ParkingLotService.GetAllocatedSlots()
		case "status_free":
			pl.ParkingLotService.GetFreeSlots()
		case "exit":
			run = false
		}
	}
}
