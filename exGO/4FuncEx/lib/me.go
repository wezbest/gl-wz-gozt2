/*
Actual solution which will be called in main
---
//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

*/

package lib

import (
	f "fmt"
)

// This will be the main function calling sub functions in this file

func Mover() {
	Texc("Greeting Function")
	Mgreet("Candy")

	Texc("Adding 3 numbers")
	f.Println(Madd(1, 2, 3))

	Texc("Display 1 num")
	f.Println(Mnum(69))

	Texc("Display 2 num")
	f.Println(Mtwo(69,420))

	Texc("Display addition of 3 numbers using above commands")
	Mthree()

	Texc("Code for the above solutions")
	ShowCode()
}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func Mgreet(name string) {
	f.Println("Lick whose pussy> :-", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func Mbark(talk string) string {
	return talk
}

//* Write a function to add 3 numbers together, supplied as arguments, and return answer 
func Madd(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func Mnum(n1 int) int {
	return n1
}

//* Write a function that returns any two numbers
func Mtwo(n1, n2 int) (int, int) {
	return n1, n2
}

//* Add three numbers together using any combination of the existing functions. 
//* Print the result
func Mthree() {
	sum := Madd(Mnum(69), Mnum(420), Mnum(420))
	f.Print(sum, "\n") 
}

func ShowCode() {
	Kode(`
	
func Mover() {
	Texc("Greeting Function")
	Mgreet("Candy")

	Texc("Adding 3 numbers")
	f.Println(Madd(1, 2, 3))

	Texc("Display 1 num")
	f.Println(Mnum(69))

	Texc("Display 2 num")
	f.Println(Mtwo(69,420))

	Texc("Display addition of 3 numbers using above commands")
	Mthree()

}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func Mgreet(name string) {
	f.Println("Lick whose pussy> :-", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func Mbark(talk string) string {
	return talk
}

//* Write a function to add 3 numbers together, supplied as arguments, and return answer 
func Madd(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func Mnum(n1 int) int {
	return n1
}

//* Write a function that returns any two numbers
func Mtwo(n1, n2 int) (int, int) {
	return n1, n2
}

//* Add three numbers together using any combination of the existing functions. 
//* Print the result
func Mthree() {
	sum := Madd(Mnum(69), Mnum(420), Mnum(420))
	f.Print(sum)
}

	`)
}

