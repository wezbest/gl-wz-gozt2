/*
Given Solution

//* Write a function that returns any message, and call it from within
//  fmt.Println()



//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once
*/

package lib

import (
	f "fmt"
)

func GiAll() {
	Texc("Greeting Function")
	f.Println("Hello World")

	Texc("Accept greeting as a parameter and then print")
	greet("Candy")

	Texc("Callling greet function without pararmeter")
	f.Println(hiThere())

	Texc("Declaring two numbers and then adding and printing on screen")
	a,b := twoTwos()
	answer:= addThree(five(),a,b)
	f.Println(answer)

	Texc("Code for the above solutions")
	ShowCodeGiv()
}


//* Write a function that returns any two numbers
func twoTwos() (int, int) {
	return 2, 2
}

//* Write a function that returns any number
func five() int {
	return 5
}	

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThree(a, b, c int) int {
	return a + b + c
}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	f.Println("Lick whose pussy> :-", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func hiThere() string {
	return "hi there"
}

//  This for  dumping code on screen 

func ShowCodeGiv() {
	Kode(`
	func GiAll() {
		Texc("Greeting Function")
		f.Println("Hello World")
	
		Texc("Accept greeting as a parameter and then print")
		greet("Candy")
	
		Texc("Callling greet function without pararmeter")
		f.Println(hiThere())
	
		Texc("Declaring two numbers and then adding and printing on screen")
		a,b := twoTwos()
		answer:= addThree(five(),a,b)
		f.Println(answer)
	}
	
	
	//* Write a function that returns any two numbers
	func twoTwos() (int, int) {
		return 2, 2
	}
	
	//* Write a function that returns any number
	func five() int {
		return 5
	}	
	
	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer
	func addThree(a, b, c int) int {
		return a + b + c
	}
	
	//* Write a function that accepts a person's name as a function
	//  parameter and displays a greeting to that person.
	func greet(name string) {
		f.Println("Lick whose pussy> :-", name)
	}
	
	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	func hiThere() string {
		return "hi there"
	}
	`)
}