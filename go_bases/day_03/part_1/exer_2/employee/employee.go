package employee

import (
	"fmt"
	"struct_emp/person"
)

type Employee struct {
	ID       int
	Position string
	person.Person
}

func (e Employee) PrintEmployee() {

	fmt.Printf("%v", e)
}
