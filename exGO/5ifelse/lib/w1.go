/*
Main work for 16_Demo_if
*/

package lib

import (
	f "fmt"
)

// This average function is is being used below
func average(a,b,c int) float32 {
	// Converting the sum of scores into a float32 
	return float32((a+b+c)/3)
}

//  Main function which will call the above
func M1() {
	quiz1, quiz2, quiz3 := 9,7,8 

	if quiz1 > quiz2 {
		f.Println("Quiz1 morethan Quiz2")
	} else if quiz1 < quiz2 {
		f.Println("Quiz2 morethan Quiz1")
	} else {
		f.Println("Quiz1 and Quiz2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		av := average(quiz1, quiz2, quiz3)
		f.Println("pass", "Score:-", av)
	} else {
		f.Println("Fucked")
	}

	showCode()
}

func showCode() {
	Kode(`
	---------------------------------------------------
	// This average function is is being used below
	func average(a,b,c int) float32 {
	// Converting the sum of scores into a float32 
	return float32((a+b+c)/3)
	}

	//  Main function which will call the above
	func M1() {
	quiz1, quiz2, quiz3 := 9,7,8 

	if quiz1 > quiz2 {
		f.Println("Quiz1 morethan Quiz2")
	} else if quiz1 < quiz2 {
		f.Println("Quiz2 morethan Quiz1")
	} else {
		f.Println("Quiz1 and Quiz2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		f.Println("pass")
	} else {
		f.Println("Fucked")
	}
}

	---------------------------------------------------
	
	`)
}