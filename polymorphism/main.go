package main

import "fmt"

// interface to implement polymorphism
type custom interface {
	nameLen() int
}

// polymorphic function to get name length
func getNameLen(c custom) int {
	return c.nameLen()
}

type person struct {
	personName string
	personAge  int
}

// nameLen impl for person type
func (p person) nameLen() int {
	fmt.Println("nameLen called with person type")
	return (len(p.personName))
}

type day struct {
	dayName string
}

// nameLen impl for day type
func (d day) nameLen() int {
	fmt.Println("nameLen called with day type")
	return (len(d.dayName))
}

func main() {

	// Object of type person
	p := person{
		personName: "Nasir",
		personAge:  25,
	}

	// Object of type day
	var d day
	d.dayName = "Wed"

	fmt.Println("personNameLen ::", getNameLen(p))
	fmt.Println("dayNameLen ::", getNameLen(d))
}
