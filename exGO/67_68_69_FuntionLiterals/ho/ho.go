/*
Function Literals / Closures
*/

package ho

import (
	"fmt"
)

func Meho() {
	MainOne()
}

// First adding two numbers

func add(lhs, rhs int) int {
	return lhs + rhs
}

// Take function as a parameter
func compute(lhs, rhs int,
	op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func MainOne() {

	// This is calling the function above
	fmt.Println("2 + 3 = ", compute(2, 3, add))

	//inline function as parameter
	fmt.Println("10 - 2 = ", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	// note this is disclosure since we are assigning function to variables
	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v =", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 3, mul))

}
