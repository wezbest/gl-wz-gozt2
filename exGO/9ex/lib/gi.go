/*
Given solution :
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides
*/

package lib

import "fmt"

// 1. Defining the coordinate struct here
type Coordinate struct {
	x, y int
}

// 2. Defining another struct called Rectangle which uses the elements of the coordinate struct above
type Rectangle3 struct {
	a Coordinate
	b Coordinate
}

// 3. Setting the width of the rectangle
func Wr3(rect Rectangle3) int {
	return (rect.b.x - rect.a.x)
}

// 4. Setting the Length of the rectangle
func Lr3(rect Rectangle3) int {
	return (rect.a.y - rect.b.y)
}

// 5. Writing the function to calculate the area
func Area3(rect Rectangle3) int {
	return Lr3(rect) * Wr3(rect)
}

// 6 Writing the function to calculate the perimeter
func Peri3(rect Rectangle3) int {
	return (Wr3(rect) * 2) + (Lr3(rect) * 2)
}

// 7. For printing the area and perimeter
func Brint(rect Rectangle3) {
	fmt.Println("Area is : ", Area3(rect))
	fmt.Println("Perimeter is : ", Peri3(rect))
}

//8 Printing the results to the terminal

func Gi() {
	Texc("Area and perimeter with coordinates set - 0,7 - 10,0")
	rect := Rectangle3{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	Brint(rect)

	Texc("Doubling the coordinates - Note the printing is being with its own function")
	rect.a.y *= 2
	rect.b.x *= 2
	Brint(rect)
}
