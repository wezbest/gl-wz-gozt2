/*
//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models


*/

package ex

import (
	f "fmt"
)

func Given() {

	// Create vehicles these will be sent to the sendToLift function
	car := Car("Sporty")
	truck := Truck("Road Devourer")
	moto := Moto("Cruzer")

	// Call the sendToLift function
	sendToLift(car)
	sendToLift(truck)
	sendToLift(moto)

}

/*
Making constants for the types
- Here the number from the iota function will be used as the index
*/

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

// Create type alias to implemen the lifts above

type Lift int

// Interface

type LiftPicker interface {
	PickLift() Lift
}

// Create vehicle type aliases
type Moto string
type Car string
type Truck string

func (m Moto) String() string {
	return f.Sprintf("Moto: %v", string(m))
}

func (c Car) String() string {
	return f.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return f.Sprintf("Truck: %v", string(t))
}

// Pick Lift interface for each type implementation

func (m Moto) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

// Single Function that can be used on all vehicles

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		f.Printf("Sending %v to small lift\n", p)
	case StandardLift:
		f.Printf("Sending %v to standard lift\n", p)
	case LargeLift:
		f.Printf("Sending %v to large lift\n", p)
	}
}
