/*
My Solution to this ex
*/

package lib

import "strconv"

// Struct for holding the products detals

type Prod struct {
	name string
	cost int
}

// Make an array for the shopping

var SL []Prod

// Making shopping list

func ShopL() {
	SL := [...]Prod{
		{name: "Apple", cost: 100},
		{name: "Banana", cost: 200},
		{name: "Mango", cost: 300},
	}
	printSL(SL)

}

// Function for adding the stuff in the array

func SumCost(SL [3]Prod) int {
	var sum int
	for i := 0; i < len(SL); i++ {
		sli := SL[i]
		sum += sli.cost
	}
	return sum
}

// Make function for printing

func printSL(SL [3]Prod) {
	Para("last Item on list is: " + SL[len(SL)-1].name)
	Para("Total number of items is: " + strconv.Itoa(len(SL)))
	Para("Total cost of items is: " + strconv.Itoa(SumCost(SL)))
}
