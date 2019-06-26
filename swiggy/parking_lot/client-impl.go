package parking_lot

import (
	"fmt"
	"sort"
	"sync"
)

type ParkingLotImpl struct{}

var (
	ParkingLotService IParkingLot
	availableSlots    []bool
	parkedSlots       map[int]Car
	lock              sync.Mutex
)

func init() {
	ParkingLotService = ParkingLotImpl{}
	parkedSlots = make(map[int]Car, 0)
}

//GetNextSlot returns the next available slot in the
//parking lot. Returns -1 if no slot is available.
func (p ParkingLotImpl) GetNextSlot() int {
	for key, val := range availableSlots {
		if val {
			return key
		}
	}
	return -1
}

//Create function creates the parking lot with n slots.
func (p ParkingLotImpl) Create(n int) {
	availableSlots = make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		availableSlots[i] = true
	}
	fmt.Printf("\nCreated a parking lot with %d slots", n)
}

//Park function parks the car in the next available slot.
//Returns the slot number when parked successfully else returns -1.
func (p ParkingLotImpl) Park(c Car) int {
	//NOTE:lock is used to handle the concurrent requests to park the car.
	lock.Lock()
	defer lock.Unlock()
	slot := p.GetNextSlot()
	if slot == -1 {
		fmt.Printf("Sorry, parking lot is full")
		return -1
	}
	parkedSlots[slot] = c
	availableSlots[slot] = false
	fmt.Printf("\nAllocated slot number: %d", slot)
	return slot
}

//Vacate function vacates the given slot and return true in case of success.
//Returns false if slot number is invalid.
func (p ParkingLotImpl) Vacate(n int) bool {
	if _, ok := parkedSlots[n]; !ok || n <= 0 || n >= len(availableSlots) {
		fmt.Println("Given slot is already vacant || Given slot number is invalid")
		return false
	}
	delete(parkedSlots, n)
	availableSlots[n] = true
	fmt.Printf("\nSlot number %d is freed", n)
	return true
}

//GetAllocatedSlots will print and
//return the list of allocated slots in parking lot
func (p ParkingLotImpl) GetAllocatedSlots() map[int]Car {
	if len(parkedSlots) > 0 {
		fmt.Println("Slot No \t Registration \t Color")
	}
	var keys []int
	for key := range parkedSlots {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%d \t %s \t %s\n", key, parkedSlots[key].RegistrationNumber, parkedSlots[key].Color)
	}
	return parkedSlots
}

//GetFree slots will print and
// return the available slots in parking lot
func (p ParkingLotImpl) GetFreeSlots() []int {
	freeSlots := make([]int, 0)
	if len(availableSlots) > 0 {
		fmt.Println("Slots")
	}
	for key, val := range availableSlots {
		if val {
			fmt.Println(key)
			freeSlots = append(freeSlots, key)
		}
	}
	return freeSlots
}
