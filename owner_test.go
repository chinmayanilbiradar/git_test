package parking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidOwner(t *testing.T) {
	want := owner{name: "John"}
	got, _ := NewOwner("John")

	assert.Equal(t, *got, want)
}

func TestOwnerWithNoName(t *testing.T) {
	_, err := NewOwner("")

	assert.NotNil(t, err)
}

func TestRegisterNewValidOwner(t *testing.T) {
	newOwner, _ := NewOwner("John")
	newLot, _ := NewParkingLot(3)
	newLot.Register(newOwner)

	assert.Contains(t, newLot.GetObservers(), newOwner)

}

