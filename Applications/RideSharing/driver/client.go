package driver

type IDriver interface {
	GetAvailableDrivers() []Driver
	StopOffering(int64)
	CanOffer(int64, int64, int64) bool
}
