/*
Structures Excercises
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
*/

package lib

import f "fmt"

// Defining the structure

type Rectangle struct {
	len int
	Hei int
}

func M1f() {
	// Set the dimensions of the rectangle

	R1 := Rectangle{
		len: 20,
		Hei: 10,
	}

	// Area and perimenter functions
	Area := R1.len * R1.Hei
	Perimeter := 2 * (R1.len + R1.Hei)

	Texc("Part1 :- Do length and height")

	T2("Rectangle Struc Details :")
	T1p(R1)

	T1("Area is (length X Height): ", f.Sprint(Area))
	T1("Perimeter is (2*(length + height)): ", f.Sprint(Perimeter))

	// Doubling the size of the rectangle 
	Texc("Part2 :- Do double the length and height")

	R2 := Rectangle {
		len: R1.len * 2,
		Hei: R1.Hei * 2,
	}

	Area2 := R2.len * R2.Hei
	Perimeter2 := 2 * (R2.len + R2.Hei)

	T2("Rectangle Struc Details :")
	T1p(R2)
	T1("Area is (length X Height): ", f.Sprint(Area2))
	T1("Perimeter is (2*(length + height)): ", f.Sprint(Perimeter2))

}
