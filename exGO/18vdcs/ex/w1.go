/*
My attempt at exercise

*/

package ex

import (
	f "fmt"

	C "github.com/fatih/color"
)

func W1func() {

	// Color definitions
	cre := C.New(C.FgRed).SprintFunc()
	cbl := C.New(C.FgBlue).SprintFunc()
	// ---

	// Declaring two slices so that they can be appended
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6}

	// Apppending function appends the two slices together 
	all := append(a, b...)

	f.Println("Appended Slices", cre(all))

	SumAlll := sum(all...)
	f.Println("Sum of Above", cbl(SumAlll))

}

// Actual variadics being written

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
