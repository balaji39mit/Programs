package parking_lot

type IParkingLot interface {
	//Create function creates the parking lot with n slots.
	Create(n int)
	//Park function parks the car in the next available slot.
	//Returns the slot number when parked successfully else returns -1.
	Park(car Car) int
	//Vacate function vacates the given slot and return true in case of success.
	//Returns false if slot number is invalid.
	Vacate(n int) bool
	//GetAllocatedSlots will print and
	//return the list of allocated slots in parking lot
	GetAllocatedSlots() map[int]Car
	//GetFree slots will print and
	//return the available slots in parking lot
	GetFreeSlots() []int
}
