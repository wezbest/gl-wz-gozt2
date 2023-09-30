/*
Com1 - This file is your efforts to write the excercise on your own

---
This command will be printing out some colorful texts with some choices

*/

package lib

import (
	"bufio"
	f "fmt"
	"os"
)

// These are set of command which are going to be put as constants

const (
	fu = "Fuck"
	su = "Suck"
	ra = "Rape"
)

// Actual Function

func Mc1() {

	// This is the scanner object capturing the stdinput
	scanner := bufio.NewScanner(os.Stdin)

	// We want to also capture the number of lines and commands which are typed
	numLines := 0
	numComs := 0

	f.Println("Press CTRL+D to exit at any time" + "\n\n")
	// Now to enter into the function
	for scanner.Scan() {
		// First checking if CTRL+D is pressed
		if scanner.Err() != nil {
		    break
		} else {
			// Extracting the text after trimming 
			text :=
		}
	}

}
