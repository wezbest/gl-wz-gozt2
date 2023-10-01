/*
Demo structures  v26
*/

package lib

import (
	f "fmt"
)

// Defining the structure

type Passenger struct {
	Name    string
	Tick    int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}

func F1() {

	// One line instantiation
	FartMistress := Passenger{"StinkMistress", 0, false}
	f.Println(FartMistress)
	T1(FartMistress)

	// Var definition of struts
	var (
		angy = Passenger{Name: "angy", Tick: 69}
		ikei = Passenger{Name: "ikai", Tick: 68}
	)
	Texc("Here using var to define structures")
	T1(angy)
	T1(ikei)
	f.Println(angy, ikei)

	// Another method of the above
	var cina Passenger
	cina.Name = "ony"
	cina.Tick = 69
	Texc("Here using var to define structures, and adding data afterwards")
	T1(cina)
	f.Println(cina)

	// Update fields
	angy.Boarded = true
	ikei.Boarded = true
	Texc("using IF to check status , after manually updating it")
	if angy.Boarded {
		f.Println("angy boarded")
	}
	if ikei.Boarded {
		f.Println("ikei boarded")
	}

	cina.Boarded = true
	bus := Bus{angy}
	Texc("Using the bust strut earlier to check if person is in front seat")
	f.Println(bus)
	f.Println(bus.FrontSeat.Name, "is in front seat")
}
