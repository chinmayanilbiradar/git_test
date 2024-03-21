package parking

import "fmt"

type Observer interface {
	update()
}

type owner struct {
	name string
}

func NewOwner(ownername string) (*owner, error) {
	if ownername == "" {
		return nil, fmt.Errorf("owner cannot have no name")
	}
	return &owner{name: ownername}, nil
}
func (o owner) update() {
	fmt.Println("Parking Lot is full")
}
