/*
Given solution to the ex
*/

package lib

import (
	f "fmt"
)

// Crate product  structure with price and a name

type Product struct {
	price int
	name  string
}

// Print stats function
func printStats(list [4]Product) {
	var cost, totalItems int

	// Making an iterator to add all the items in the list
	for i := 0; i < len(list); i++ {
		item := list[i]    // Making a copy for concurrency
		cost += item.price // Here Adding the total of the cost of items
		if item.name != "" {
			totalItems += 1
		}
	}

	// Printing the stuff
	f.Println("Last item on list is: ", list[totalItems-1])
	f.Println("Total items", totalItems)
	f.Println("Total Cost ", cost)

}

func Gi() {
	shoppingList := [4]Product{
		{1, "banana"},
		{2, "apple"},
		{3, "mango"},
	}

	printStats(shoppingList)

	/* 
	This code is being used to add to the array 
	- Since array starts at index 0 , adding a 3rd tiem will make it the fourth 
	- This being done in code below , where the fourth item is 4, Bread
	*/
	shoppingList[3] = Product{4, "Bread"}

	printStats(shoppingList)
}
