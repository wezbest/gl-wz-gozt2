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

// This is going to be the computer function which also has a type embedding from above
type op func(lhs, rhs int) int

func compute2(lhs, rhs int, op op) int {
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
	// Anonymous function - Closure
	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v =", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 3, mul))

	// This is adding the same compute function with the new one
	fmt.Println("This is second compute with type embedding :",compute2(3, 3, mul))
}
