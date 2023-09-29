/*
Exercise my attempt
*/

package lib

import (
	f "fmt"
)

// Solution with creating the slice
// -from the given solution , you can define the type of any variable as sa string , sot tha tis implemented in the given solution
func Mys() {
	Texc("Original Parts")
	asml := []string{"Part1", "Part2", "Part3"}
	myPr("Original Parts", asml)

	Texc("Added Parts")
	asml = append(asml, "Part4", "Part5")
	myPr("Added Parts", asml)

	Texc("Newly Added Parts")
	asml = asml[3:]
	myPr("Newly Added Parts", asml)
}

// Helper function to print out the slice
func myPr(title string, slice []string) {
	f.Println()
	f.Println("---", title, "---")

	// Writing a loop iterate and print out the slice
	for i := 0; i < len(slice); i++ {
		element := slice[i] // This is making copy to prevent cooncurrency issues later on
		f.Println(element)
	}
}
