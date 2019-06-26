package main

type Position string

const (
	CityManager        Position = "CM"
	FleetManager                = "FM"
	DeliveryExecutives          = "DE"
)

type Rating uint8

const (
	Outstanding Rating = 5
	Excellent          = 4
	Good               = 3
	Bad                = 2
	Poor               = 1
)

var BonusPercentage map[Rating]float64

func init() {
	BonusPercentage = make(map[Rating]float64)
	BonusPercentage[Outstanding] = 120
	BonusPercentage[Excellent] = 110
	BonusPercentage[Good] = 100
	BonusPercentage[Bad] = 90
	BonusPercentage[Poor] = 80
}

const (
	MAXFLOAT float64 = 1000000000
)
