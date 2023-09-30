/*
Com1 - This file is your efforts to write the excercise on your own

---
This command will be printing out some colorful texts with some choices

*/

package lib

import (
	"bufio"
	"fmt"
	f "fmt"
	"os"
	"strings"
)

// These are set of command which are going to be put as constants

const (
	fu = "Fuck"
	su = "Suck"
	ra = "Rape"
	ex = "Exit"
)

// Actual Function

func Mc1() {

	// Dcclaring the colors here

	// This is the scanner object capturing the stdinput
	scanner := bufio.NewScanner(os.Stdin)

	// We want to also capture the number of lines and commands which are typed
	numLines := 0
	numComs := 0

	f.Println("Press CTRL+D or type \"EXIT\" to exit at any time" + "\n\n")
	// Now to enter into the function
	for scanner.Scan() {
		// First checking if CTRL+D is pressed
		if scanner.Err() != nil {
			break
		} else {
			// Extracting the text after trimming
			text := strings.TrimSpace(scanner.Text())

			// Now writing the switch cases fir the commands that are available
			switch text {
			case fu:
				// Code for "Fuck" command
				T1("Fuckyou")
				numComs++
			case su:
				// Code for "Suck" command
				T2("Suck Ass")
				numComs++
			case ra:
				// Code for "Rape" command
				T3("Get Raped")
				numComs++
			case ex:
				// Code for "Exit" command
				os.Exit(1)
			default:
				// Code for unrecognized command
				f.Println("😡")
			}
			if text != "" {
				numLines++
			}
		}
	}
	fmt.Printf("Entered Commands %v \n", numComs)
	fmt.Printf("Entered Lines %v \n", numLines)

}
