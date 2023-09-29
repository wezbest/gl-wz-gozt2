/*


 ██████╗  ███████╗ ███╗   ███╗  ██████╗     
 ██╔══██╗ ██╔════╝ ████╗ ████║ ██╔═══██╗    
 ██║  ██║ █████╗   ██╔████╔██║ ██║   ██║    
 ██║  ██║ ██╔══╝   ██║╚██╔╝██║ ██║   ██║    
 ██████╔╝ ███████╗ ██║ ╚═╝ ██║ ╚██████╔╝    
 ╚═════╝  ╚══════╝ ╚═╝     ╚═╝  ╚═════╝     

██╗   ██╗  █████╗  ██████╗  ██╗  █████╗  ██████╗  ██╗      ███████╗ ███████╗    
██║   ██║ ██╔══██╗ ██╔══██╗ ██║ ██╔══██╗ ██╔══██╗ ██║      ██╔════╝ ██╔════╝    
██║   ██║ ███████║ ██████╔╝ ██║ ███████║ ██████╔╝ ██║      █████╗   ███████╗    
╚██╗ ██╔╝ ██╔══██║ ██╔══██╗ ██║ ██╔══██║ ██╔══██╗ ██║      ██╔══╝   ╚════██║    
 ╚████╔╝  ██║  ██║ ██║  ██║ ██║ ██║  ██║ ██████╔╝ ███████╗ ███████╗ ███████║    
  ╚═══╝   ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚═╝ ╚═╝  ╚═╝ ╚═════╝  ╚══════╝ ╚══════╝ ╚══════╝    

██╗    ██╗  ██████╗  ██████╗  ██╗  ██╗
██║    ██║ ██╔═══██╗ ██╔══██╗ ██║ ██╔╝
██║ █╗ ██║ ██║   ██║ ██████╔╝ █████╔╝ 
██║███╗██║ ██║   ██║ ██╔══██╗ ██╔═██╗ 
╚███╔███╔╝ ╚██████╔╝ ██║  ██║ ██║  ██╗
 ╚══╝╚══╝   ╚═════╝  ╚═╝  ╚═╝ ╚═╝  ╚═╝
Actual lessons are in here which are then called in main.go
*/

package lib

import (
	f "fmt" // Format package from standard library
)

// These are functions , will make a separate one for each here
func Fu1() {
	var myName = "BootyAddict"
	f.Println("My fetish is", myName, myName)
	Kode(`
	func Fu1() {
		var myName = "BootyAddict"
		f.Println("My fetish is", myName, myName)
	} 

	// In this function using var to define varibale which is not needed
	`)
} 

// Type annotations
func Fu2() {
	var fetish string = "Sniffer"
	f.Println("My fetish is", fetish, fetish)
	Kode(`
	func Fu2() {
		var fetish string = "Sniffer"
		f.Println("My fetish is", fetish, fetish)
	}

	// Here youc an see we are defining [Tyoe] for the variable which is being then used
	`)
}

// assign and create method 
func Fu3() {
	userName := "PantySniffer"
	f.Println("My usrename is", userName, userName)
	Kode(`
	func Fu3() {
		userName := "PantySniffer"
		f.Println("My usrename is", userName, userName)
	}

	// Here is a direct assignment
	`)
}

// Uniintialized variable 
func Fu4() {
	var sum int 
	f.Println("The sum is = ", sum)
	Kode(`
	func Fu4() {
		var sum int 
		f.Println("The sum is = ", sum)
	}

	// This is an uninitialized variable, the default value will be 0
	`)
}

// Complex assigment
func Fu5() {
	big, booty := 6,9 
	f.Println("Big = ", big, "Booty = ", booty)
	Kode(`
	func Fu5() {
		big, booty := 6,9 
		f.Println("Big = ", big, "Booty = ", booty)
	}

	// This is a mutiple assignments with inferred types from go
	`)
}

//OK error idiom 
func Fu6() {
	sml, boty := 6,9
	byg, boty := 6,9
	f.Println("Big = ", sml, "Booty = ", boty, "This written to remove error", byg)
	Kode(`
	func Fu6() {
		sml, boty := 6,9
		byg, boty := 6,9
		f.Println("Big = ", sml, "Booty = ", boty, "This written to remove error", byg)
	}
	// This is illustrating how same varibale can be assigned two times in the same code block
	// And it will have the new value
	`)
}

// Reassign variables 
func Fu7() {
	var sum int
	booty, titty := 6,9
	sum = booty + titty
	f.Println("The sum is = ", sum)
	Kode(`
	func Fu7() {
		var sum int
		booty, titty := 6,9
		sum = booty + titty
		f.Println("The sum is = ", sum)
	}

	// Illustration of an unused variable which can be used at a later time 
	// Note in the actual tutorial all of this was being done inside on function , but you are making 
	// Seperate functions for you own practice 
	`)

}

// 8 - Block Asignment 

func Fu8() {
	var (
		pantyStyle = "Thong"
		pantySize = "InsideAsss"
	)
	f.Println("Mistress Style:",pantyStyle, pantySize)

	Kode(`
	func Fu8() {
		var (
			pantyStyle = "Thong"
			pantySize = "InsideAsss"
		)
		f.Println("Mistress Style:",pantyStyle, pantySize)
	}

	// From the last tut sweaty mistress asked you to do it the wrong way , and you saw this from examinig
	// Code block, from another reppo -The lipgloss repo which is also what you are using to print all the 
	// color stuffs here
	`)
}

// 9 - Compound assignment with ignoring a variable 

func Fu9() {

	lick, sniff, _ := "Pussu", "Ass", "fart"
	f.Println(lick, sniff)

	Kode(`
	func Fu9() {

		lick, sniff, _ := "Pussu", "Ass", "fart"
		f.Println(lick, sniff)
	}

	// Note the underscore "_" , this is for ignoring the variable which you had used previously in 
	// For loops
	`)
}

