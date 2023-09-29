/*
V24 - You attempt at solution
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
*/

package lib

import (
	f "fmt"
	r "math/rand"
	ti "time"
)

// Main runner

func Mmaina() {
	Mplay4()

}

// Dice roll

// This is the function for rolling the dice
func Mdic() {
	d1 := r.Intn(6) + 1
	d2 := r.Intn(6) + 1
	dtotal := d1 + d2
	T1("Dice 1: " + f.Sprint(d1))
	T1("Dice 2: " + f.Sprint(d2))
	T2("Dice Total:- "+f.Sprint(dtotal), "", "")
}

// This is rolling the dice 10 times to see what happens
func Mplay() {
	for i := 0; i < 10; i++ {
		T3("\n" + "Roll Number: " + f.Sprint(i))
		Mdic()
	}
}

// This function will check for the given conditions
func Mplay2() {

	// Setting up the dices here and calculating the total
	d1 := r.Intn(6) + 1
	d2 := r.Intn(6) + 1
	dtotal := d1 + d2

	// Printing what is the dice
	T1("Dice 1: " + f.Sprint(d1))
	T1("Dice 2: " + f.Sprint(d2))
	T2("Dice Total:- "+f.Sprint(dtotal), "", "")

	// Checking for the given conditions
	if dtotal == 2 && d1 == 2 && d2 == 2 {
		T3("Snake Eyes")
	} else if dtotal == 7 {
		T3("Lucky 7")
	} else if dtotal%2 == 0 {
		T1("Even")
	} else {
		T2("Odd", "", "")
	}
}

// Now lets run the dice 20 times and see what happens
func Mplay3() {
	for i := 0; i < 20; i++ {
		T3("\n" + "Roll Number: " + f.Sprint(i))
		Mplay2()
	}
}

/*
Now using the switch statment to check for the conditions insted of If above
- Why because performance wise switch is faster when u readingtoninClaudeArmpitSniffing
- Here we are also adding a random seed, based on the current unix time
*/

func Mplay4() {

	// Adding this line to ensure additional randomness
	t := r.New(r.NewSource(ti.Now().UnixNano()))
	f.Println("Random seed is :", t)

	d1 := t.Intn(6) + 1
	d2 := t.Intn(6) + 1
	dtotal := d1 + d2

	T1("Dice 1: " + f.Sprint(d1))
	T1("Dice 2: " + f.Sprint(d2))
	T2("Dice Total:- "+f.Sprint(dtotal), "", "")

	// Note this section fully written by codeiumSquirtDrink
	/*
		Explanation of the logic here  :
		1. dtotal is the sum of the two dice
		2. d1 and d2 are the dice
		3. d1 + d2 is the total
		4. if dtotal == 2 and d1 == 2 and d2 == 2 then we have snake eyes
		5. if dtotal == 7 then we have lucky 7
		6. if dtotal is even then we have even
		7. if dtotal is odd then we have odd
		- So if the total is neither snake eyes.
	*/
	switch {
	case dtotal == 2 && d1 == 2 && d2 == 2:
		T3("Snake Eyes")
	case dtotal == 7:
		T3("Lucky 7")
	case dtotal%2 == 0:
		T1("Even")
	default:
		T2("Odd", "", "")
	}
}

// Testing some things form above code 

func Mplay6() {
	 d1, d2 := r.Intn(6)+1, r.Intn(6)+1
	 f.Println(d1, d2)
}