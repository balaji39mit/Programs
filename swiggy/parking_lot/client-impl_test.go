package parking_lot

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var (
	mockParkingLotService IParkingLot
)

func init() {
	mockParkingLotService = ParkingLotImpl{}
}

func TestCreate(t *testing.T) {
	assert.Nil(t, availableSlots)
	mockParkingLotService.Create(6)
	assert.NotNil(t, availableSlots)
	assert.Equal(t, availableSlots, []bool{false, true, true, true, true, true, true}, fmt.Sprintf("Create is failing.."))
}

func TestPark(t *testing.T) {

	testData := []struct {
		name   string
		input  Car
		output int
	}{
		{
			name:   "Success-ParkFirst",
			input:  Car{RegistrationNumber: "KA-01-HH-1234", Color: "White"},
			output: 1,
		},
		{
			name:   "Success-ParkSecond",
			input:  Car{RegistrationNumber: "KA-01-HH-9999", Color: "White"},
			output: 2,
		},
		{
			name:   "Success-ParkThird",
			input:  Car{RegistrationNumber: "KA-01-HH-0001", Color: "White"},
			output: 3,
		},
		{
			name:   "Success-ParkFourth",
			input:  Car{RegistrationNumber: "KA-01-HH-9999", Color: "White"},
			output: 4,
		},
		{
			name:   "Success-ParkFifth",
			input:  Car{RegistrationNumber: "KA-01-HH-1234", Color: "White"},
			output: 5,
		},
		{
			name:   "Success-ParkSixth",
			input:  Car{RegistrationNumber: "KA-01-HH-9999", Color: "White"},
			output: 6,
		},
		{
			name:   "Failure",
			input:  Car{RegistrationNumber: "3", Color: "RED"},
			output: -1,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res := mockParkingLotService.Park(td.input)
			assert.Equal(t, td.output, res, fmt.Sprintf("Expected : %v --- Actual : %v", td.output, res))
		})
	}
}

func TestVacate(t *testing.T) {
	testData := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "Success",
			input:  4,
			output: true,
		},
		{
			name:   "Failure-MaxLength",
			input:  10,
			output: false,
		},
		{
			name:   "Failure-MinLength",
			input:  -3,
			output: false,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res := mockParkingLotService.Vacate(td.input)
			assert.Equal(t, res, td.output, fmt.Sprintf("Expectted : %v --- Actual : %v", td.output, res))
		})
	}
}

func TestGetAllocatedSlots(t *testing.T) {
	testData := []struct {
		name   string
		output map[int]Car
	}{
		{
			name: "Success",
			output: map[int]Car{
				1: Car{RegistrationNumber: "KA-01-HH-1234", Color: "White"},
				2: {RegistrationNumber: "KA-01-HH-9999", Color: "White"},
				3: {RegistrationNumber: "KA-01-HH-0001", Color: "White"},
				5: {RegistrationNumber: "KA-01-HH-1234", Color: "White"},
				6: {RegistrationNumber: "KA-01-HH-9999", Color: "White"},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res := mockParkingLotService.GetAllocatedSlots()
			assert.Equal(t, res, td.output, fmt.Sprintf("Expectted : %v --- Actual : %v", td.output, res))
		})
	}

}

func TestGetFreeSlots(t *testing.T) {
	testData := []struct {
		name   string
		output []int
	}{
		{
			name:   "Success",
			output: []int{1, 2},
		},
	}

	mockParkingLotService.Create(2)

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res := mockParkingLotService.GetFreeSlots()
			assert.Equal(t, res, td.output, fmt.Sprintf("Expected : %v --- Actual : %v", td.output, res))
		})
	}

}

func Test(t *testing.T) {
	TestCreate(t)
	TestPark(t)
	TestVacate(t)
	TestGetAllocatedSlots(t)
}
