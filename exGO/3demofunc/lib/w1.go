/*
Actual work for demo files 
*/

package lib 

import (
	f "fmt"
)

// Demo Function 1  - Doubles number

func F1dbl(x int) int {
	return x + x 
}

// Demo Function 2 - Add Number 
func F2add(lhs, rhs int) int {
	return lhs + rhs
}

func F3greet() {
	f.Println("Sniff her")
}

func F4ac() {
	Kode(`
	// Above results from the following code
	
	func F1dbl(x int) int {
		return x + x 
	}
	
	// Demo Function 2 - Add Number 
	func F2add(lhs, rhs int) int {
		return lhs + rhs
	}
	
	func F3greet() {
		f.Println("Sniff her")
	}

	// Following is the code in main.go 
	func main() {
		l.ColoBanner("Demo Functions work")
		
		l.F3greet()
	
		dozen := l.F1dbl(69)
		f.Println(dozen)
	
		bakerDozen := l.F2add(dozen,1)
		f.Println("baker Dozen: ", bakerDozen)
	
		// Nesting functions
		anotherBakerDozen := l.F2add(l.F1dbl(6),1)
		f.Println("another baker Dozen: ", anotherBakerDozen)
	
		// Printing source 
		l.F4ac()
	}

	`)
}