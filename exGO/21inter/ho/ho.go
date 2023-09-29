/*
Demo Exercise
Interface for preparation of food , one interface will be called for food prep
*/

package ho

import (
	f "fmt"
	L "rape/lib"
)

func Meho() {
	// Main Function is being called here
	dishes := []Preparer{
		Chicken("Hot Wings"),
		Salad("Greek Salad"),
	}
	prepareDishes(dishes)
}

// Create interface
// Adding - er as sa suffix will make it an interface

type Preparer interface {
	PrepareDish()
}

// Now above function will have two dishes here

type Chicken string
type Salad string

// Function is for preparing the chicken
func (c Chicken) PrepareDish() {
	L.T1("Cook Chicken")

}

// Function is for preparing the salad
func (s Salad) PrepareDish() {
	L.T2("Chop Salad")
	L.T2("Add Dressing")
}

// Function prepare dishes

func prepareDishes(dishes []Preparer) {
	f.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		f.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	f.Println()
}
