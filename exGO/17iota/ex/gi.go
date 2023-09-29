/*
Given solution 
*/

package ex

import (
	f "fmt"
)

func Given() {
	// This is the main function which needs to be altered 

	add := operation(Add)
	f.Println(add.calculate(2, 2)) // = 4

	sub := operation(Sub)
	f.Println(sub.calculate(10, 3)) // = 7

	mul := operation(Mul)
	f.Println(mul.calculate(3, 3)) // = 9

	div := operation(Div)
	f.Println(div.calculate(100, 2)) // = 50

}

// Defining the iota 
/*
Using this method is also called the iota enumeration pattern
*/

const (
	Add = iota
	Sub 
	Mul
	Div
)

/* 
Tyoe alias for cosntants , meaning the above 
- The iota above will become integers when using the operation int type
- The operation int type itself is a receiver function , so it will have  function like
> operation.add 
> operation.sub
> operation.mul
> operation.div

*/

type operation int  // Defining one variable for storing operation 


// Writing the add function 
// Here the iota in the constants is used for assigning number to the cases 
func (op operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Sub:
		return lhs - rhs
	case Mul:
		return lhs * rhs
	case Div:
		return lhs / rhs
	}
	panic("Bastard what the fuck are you doing ? ")
}