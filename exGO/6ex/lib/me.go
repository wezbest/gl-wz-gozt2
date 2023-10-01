/*
This is going to be your work on this file
*/

package lib

import (
	f "fmt"
)

// Function for getting the age

func Age() int {
	return 100
}

func M1F() {
	switch ag := Age(); {
	case ag == 0:
		f.Println("newborn", "Age is:", ag)
	case ag >= 1 && ag <= 3:
		f.Println("toddler", "Age is:", ag)
	case ag >= 4 && ag <= 12:
		f.Println("child", "Age is:", ag)
	case ag >= 13 && ag <= 18:
		f.Println("teenager" ,"Age is:", ag)
	default:
		f.Println("adult", "Age is:", ag)
	}

	ShowCode()
}

func ShowCode() {
	Kode(`
	
	func Age() int {
		return 0
	}
	
	func M1F() {
		switch ag := Age(); {
		case ag == 0:
			f.Println("newborn")
		case ag >= 1 && ag <= 3:
			f.Println("toddler")
		case ag >= 4 && ag <= 12:
			f.Println("child")
		case ag >= 13 && ag <= 18:
			f.Println("teenager")
		default:
			f.Println("adult")
		}
	}

	`)
}
