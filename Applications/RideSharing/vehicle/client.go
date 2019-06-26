package vehicle

type IVehicle interface {
	CanServeRide(int64, int64) bool
	GetPrice(id int64) float64
	Display(id int64) string
}
