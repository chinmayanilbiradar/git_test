package parking

import "fmt"

type ParkingLot struct {
	parkingLotObservers []Observer
	size                int
	ListOfCars          []Car
}

type Observable interface {
	Register()
	Notify()
}

type Car struct {
	NumberPlate int
}

func NewCar(id int) *Car {
	return &Car{NumberPlate: id}
}
func NewParkingLot(size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, fmt.Errorf("cannot create parkingLot of negative/zero size")
	}
	return &ParkingLot{parkingLotObservers: make([]Observer, 0), size: size, ListOfCars: make([]Car, size)}, nil
}

func (p *ParkingLot) GetSize() int {
	return p.size
}

func (p *ParkingLot) GetObservers() []Observer {
	return p.parkingLotObservers
}

func (p *ParkingLot) Park(c Car) error {
	if p.size <= 0 {
		return fmt.Errorf("no parking available")
	}
	p.ListOfCars = append(p.ListOfCars, c)
	p.size -= 1
	return nil
}

func (p ParkingLot) checkAvailability() (bool, error) {
	if p.size <= 0 {
		return false, fmt.Errorf("no parking available")
	}
	return true, nil
}

func (p *ParkingLot) RemoveCar(c Car) (*ParkingLot, error) {
	flag := false
	for i, v := range p.ListOfCars {
		if v.NumberPlate == c.NumberPlate {
			p.ListOfCars[i] = p.ListOfCars[len(p.ListOfCars)-1]
			p.ListOfCars[len(p.ListOfCars)-1] = Car{NumberPlate: 0}
			p.ListOfCars = p.ListOfCars[:len(p.ListOfCars)-1]
			// p.ListOfCars = p.ListOfCars[:i] + p.ListOfCars[i+1:]
			flag = true
			break
		}
	}
	if flag {
		return p, nil
	}
	return nil, fmt.Errorf("Car not present in parking Lot")
}

func (p *ParkingLot) Register(observer Observer) {
	p.parkingLotObservers = append(p.parkingLotObservers, observer)
}
