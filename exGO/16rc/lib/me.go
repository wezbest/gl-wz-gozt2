/*
Tracking stats of parking places
*/

package lib

import (
	f "fmt"

	"github.com/fatih/color"
)

// Creating structures

// Represents single space
type Space struct {
	occupied bool
}

// Represets all the parking places
type ParkingLot struct {
	spaces []Space
}

// Regular function - This function indicates that a parking lot space is occupied
func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// Now create receiver function  - Same like above function but with a receiver
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// This function indicates that a space has been vacated - opposite of the above
func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func Me() {

	//Defining the colors with fatih
	cg := color.New(color.FgGreen).SprintFunc()
	cr := color.New(color.FgRed).SprintFunc()
	cb := color.New(color.FgBlue).SprintFunc()

	// This is your addition when making the parking lot
	Spacess := 3

	// Setting up the parking lots
	lot := ParkingLot{spaces: make([]Space, Spacess)} // Defining total number of spaces
	f.Println("Initial Parking places:", cg(lot))     // Printing the total number of spaces

	// Regular Function method
	lot.occupySpace(1) // Occupying space 1

	// Receiver function method here
	occupySpace(&lot, 2)

	f.Println("After occupying", cb(lot)) // Printing the total number of spaces

	//Vacate spaces
	lot.vacateSpace(2)
	f.Println("After Vacate", cr(lot)) // Printing the total number of spaces
}
