package parking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParkingLotForSize10(t *testing.T) {
	want := ParkingLot{parkingLotObservers: make([]Observer, 0), size: 10, ListOfCars: make([]Car, 10)}
	got, _ := NewParkingLot(10)

	assert.Equal(t, *got, want, "they should be equal")
}

func TestNewParkingLotforSize0(t *testing.T) {
	_, err := NewParkingLot(0)
	errCheck := "cannot create parkingLot of negative/zero size"
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), errCheck, "test should be equal")
	}
}

func TestAvailabilityForSize5(t *testing.T) {
	newLot, _ := NewParkingLot(5)
	got, err := newLot.checkAvailability()
	want := true
	if assert.Nil(t, err) {
		assert.Equal(t, got, want)
	}
}

func TestParkForParkingLotOfSize1(t *testing.T) {
	want := 0
	newLot, _ := NewParkingLot(1)
	newCar := NewCar(1)
	err := newLot.Park(*newCar)
	if assert.Nil(t, err) {
		got := newLot.size
		assert.Equal(t, got, want)
	}
}

func TestParkForParkingLotOfSize8(t *testing.T) {
	want := 7
	newLot, _ := NewParkingLot(8)
	newCar := NewCar(1)
	err := newLot.Park(*newCar)
	if assert.Nil(t, err) {
		got := newLot.size
		assert.Equal(t, got, want)
	}
}

func TestParkForParkingLotAfterParkingLotBecomesFull(t *testing.T) {
	errCheck := "no parking available"
	newLot, _ := NewParkingLot(1)
	newCar := NewCar(1)
	_ = newLot.Park(*newCar)

	err2 := newLot.Park(*newCar)

	assert.Equal(t, err2.Error(), errCheck)

}

func TestRemoveCarWhichDoesNotExistInParkingLot(t *testing.T) {
	newLot, _ := NewParkingLot(3)
	newCar := NewCar(1)

	_, err := newLot.RemoveCar(*newCar)

	assert.NotNil(t, err)
}

func TestRemoveCarWhichExistsInParkingLot(t *testing.T) {
	newLot, _ := NewParkingLot(3)
	newCar := NewCar(1)
	_ = newLot.Park(*newCar)
	want := len(newLot.ListOfCars) - 1

	got, _ := newLot.RemoveCar(*newCar)

	assert.Equal(t, len(got.ListOfCars), want)

}

func TestRemoveCarWhichIsAlreadyRemoved(t *testing.T) {
	newLot, _ := NewParkingLot(3)
	newCar := NewCar(1)
	_ = newLot.Park(*newCar)
	got, _ := newLot.RemoveCar(*newCar)

	_, err2 := got.RemoveCar(*newCar)

	assert.NotNil(t, err2)
}
