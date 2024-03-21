package parking

import "fmt"

type Cop struct {
	copName string
}

func (c Cop) Update() {
	fmt.Println("Redirect Traffic")
}
