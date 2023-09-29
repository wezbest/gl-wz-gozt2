/*
Demo Functions work 
*/

package main 

import (
	l "df/lib"
	f "fmt"
)

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