/*
 ██████╗  ███████╗ ███╗   ███╗  ██████╗      ███████╗ ██╗  ██╗  ██████╗ ███████╗ ██████╗   ██████╗ ██╗ ███████╗ ███████╗
 ██╔══██╗ ██╔════╝ ████╗ ████║ ██╔═══██╗     ██╔════╝ ╚██╗██╔╝ ██╔════╝ ██╔════╝ ██╔══██╗ ██╔════╝ ██║ ██╔════╝ ██╔════╝
 ██║  ██║ █████╗   ██╔████╔██║ ██║   ██║     █████╗    ╚███╔╝  ██║      █████╗   ██████╔╝ ██║      ██║ ███████╗ █████╗
 ██║  ██║ ██╔══╝   ██║╚██╔╝██║ ██║   ██║     ██╔══╝    ██╔██╗  ██║      ██╔══╝   ██╔══██╗ ██║      ██║ ╚════██║ ██╔══╝
 ██████╔╝ ███████╗ ██║ ╚═╝ ██║ ╚██████╔╝     ███████╗ ██╔╝ ██╗ ╚██████╗ ███████╗ ██║  ██║ ╚██████╗ ██║ ███████║ ███████╗
 ╚═════╝  ╚══════╝ ╚═╝     ╚═╝  ╚═════╝      ╚══════╝ ╚═╝  ╚═╝  ╚═════╝ ╚══════╝ ╚═╝  ╚═╝  ╚═════╝ ╚═╝ ╚══════╝ ╚══════╝

     ███████╗  ██████╗  ██╗      ██╗   ██╗ ████████╗ ██╗  ██████╗  ███╗   ██╗ ███████╗
     ██╔════╝ ██╔═══██╗ ██║      ██║   ██║ ╚══██╔══╝ ██║ ██╔═══██╗ ████╗  ██║ ██╔════╝
     ███████╗ ██║   ██║ ██║      ██║   ██║    ██║    ██║ ██║   ██║ ██╔██╗ ██║ ███████╗
     ╚════██║ ██║   ██║ ██║      ██║   ██║    ██║    ██║ ██║   ██║ ██║╚██╗██║ ╚════██║
     ███████║ ╚██████╔╝ ███████╗ ╚██████╔╝    ██║    ██║ ╚██████╔╝ ██║ ╚████║ ███████║
     ╚══════╝  ╚═════╝  ╚══════╝  ╚═════╝     ╚═╝    ╚═╝  ╚═════╝  ╚═╝  ╚═══╝ ╚══════╝

	 -------------------------------------------------------------

This will be solutions to the exceercises, which will then be called in main.go

*/

package lib

import (
	f "fmt"
)

// Following is my solution
func MySoln() {

	// 1 - Store your favorite color in a variable using the `var` keyword
	var favColor = "pink"
	f.Println("1 - Fav Pussy Color :-", favColor)

	// 2 - Store your birth year and age (in years) in two variables using compound assignment
	botyYer, botyAge := 2000, 18
	f.Println("2 - Birth Year and Age :-", botyYer, botyAge)

	// 3 - Store your first & last initials in two variables using block assignment
	var (
		firstLick = 'B'
		lastLick  = 'B'
	)
	f.Println("3 - First & Last Initials :-", firstLick, lastLick)

	//4 - Declare (but don't assign!) a variable for your age (in days), then assign it
	// on the next line by multiplying 365 with the age variable created earlier
	
	var ageInDays int 
	ageInDays = 69 * 365
	f.Println("4 - Age in Days :-", ageInDays)
	
	Kode(`

	// The following is your solution 
	func Fu1() {

		// 1 - Store your favorite color in a variable using the var keyword
		var favColor = "pink"
		f.Println("1 - Fav Pussy Color :-", favColor)
	
		// 2 - Store your birth year and age (in years) in two variables using compound assignment
		botyYer, botyAge := 2000, 18
		f.Println("2 - Birth Year and Age :-", botyYer, botyAge)
	
		// 3 - Store your first & last initials in two variables using block assignment
		var (
			firstLick = "B"
			lastLick  = "B"
		)
		f.Println("3 - First & Last Initials :-", firstLick, lastLick)
	
		//4 - Declare (but don't assign!) a variable for your age (in days), then assign it
		// on the next line by multiplying 365 with the age variable created earlier
		
		var ageInDays int 
		ageInDays = 69 * 365
		f.Println("4 - Age in Days :-", ageInDays)
	}
	`)
	

}

// Given Solution 

func GvnSoln() {

	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor = "red"
	f.Println("1 - Fav Pussy Color :-", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInyears := 2000, 18
	f.Println("Born in", birthYear, "and is", ageInyears, "years old")
	
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "B"
		lastInitial  = "B"
	)
	f.Println("1st Initia", firstInitial, "last Initial", lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier

	var ageInDays int
	ageInDays = ageInyears * 365
	f.Println("Age in Days :-", ageInDays, "days old")

	Kode(`
	// Given solution
	func GvnSoln() {

		//* Store your favorite color in a variable using the var keyword
		var favoriteColor = "red"
		f.Println("1 - Fav Pussy Color :-", favoriteColor)
	
		//* Store your birth year and age (in years) in two variables using
		//  compound assignment
		birthYear, ageInyears := 2000, 18
		f.Println("Born in", birthYear, "and is", ageInyears, "years old")
		
		//* Store your first & last initials in two variables using block assignment
		var (
			firstInitial = "B"
			lastInitial  = "B"
		)
		f.Println("1st Initia", firstInitial, "last Initial", lastInitial)
	
		//* Declare (but don't assign!) a variable for your age (in days),
		//  then assign it on the next line by multiplying 365 with the age
		// 	variable created earlier
	
		var ageInDays int
		ageInDays = ageInyears * 365
		f.Println("Age in Days :-", ageInDays, "days old")
	
	}
	`)	
}