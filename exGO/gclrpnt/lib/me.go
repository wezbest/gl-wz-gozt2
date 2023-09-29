/*
My attemtp at excercise
*/

package lib

import (
	f "fmt"
)

func Me() {
	StrPr()
}

// Testing out color print of various data structures

func StrPr() {
	// Defining and printing array
	a := []int{1, 2, 3}
	Para("Printing original array without colring and using fmt.println")
	f.Print(a, "\n")
	Para("Color pringint arrays")
	T1(a)
	f.Println("\n")

	// Defiining and printing a map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	Para("Printing original map without colring and using fmt.println")
	f.Print(m, "\n")
	Para("Color pringint maps")

}
