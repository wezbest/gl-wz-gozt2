/*
Given solution
*/

package lib

import (
	f "fmt"
)

func GiF() {
	shirt := Item{"Shirt", Active}
	Pants := Item{"Pants", Active}
	Shoes := Item{"Shoes", Active}
	Pen := Item{"Pen", Active}

	// Here the slice is being made from the struct - Which you were wondering earlier
	items := []Item{shirt, Pants, Shoes, Pen}
	f.Println("Main list", items)
	deactivate(&items[0].Tag)
	f.Println("1st Item checkedout", items)
	checkout(items)
	f.Println("checkedout", items)
}

// Creating a constant so that we can check something it active or inactive
const (
	Active   = true
	Deactive = false
)

// Create type to manange the state

type SecurityTag bool

// Create a structure to store items and their security tag state
type Item struct {
	Name string
	Tag  SecurityTag
}

// Activate function , which is creating a pointer to the tag
func activate(tag *SecurityTag) {
	*tag = Active
}

// Deactivate function , which is creating a pointer to the tag
func deactivate(tag *SecurityTag) {
	*tag = Deactive
}

// Creating the function to deactivate the tags
// Here the pointer is working on the items struct which has been made above
func checkout(items []Item) {
	f.Println("Checkint out")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].Tag)
	}
}
