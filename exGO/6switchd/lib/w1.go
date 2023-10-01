/*
19_Switch Case Demo
*/

package lib

import (
	f "fmt"
)

// Code which was included

func price() int {
	return 100
}

const (
	Economy  = 0
	Business = 1
	FirstC   = 2
)

func M1f() {
	switch p := price(); {
	case p < 2:
		f.Println("Cheap", "Price is:", p)
	case p < 10:
		f.Println("Middle", "Price is:", p)
	default:
		f.Println("Expensive Pussy", "Price is:", p)
	}

	ticket := Economy
	switch ticket {
		case Economy:
			f.Println("Economy Seating", "becoz seat is:", ticket)
		case Business:
			f.Println("Business Seating","becoz seat is:", ticket)
		case FirstC:
			f.Println("First Class","becoz seat is:", ticket)
		default:
			f.Println("NoPlane","becoz seat is:", ticket)
	}
}

func ShowCode() {
	Kode(`
	func price() int {
		return 1
	}
	
	const (
		Economy = 0 
		Business = 1 
		FirstC = 2
	)
	
	func M1f() {
		switch p:=price(); {
		case p < 2 :
			f.Println("Cheap")
		case p < 10 :
			f.Println("Middle")
		default :
			f.Println("Expensive Pussy")
		}
	}
	
	`)
}
