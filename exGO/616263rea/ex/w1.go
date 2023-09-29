/*
My attempt at exercise
*/

package ex

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Building the command here
*/

// Commands are hello and bye and those first will be constants

const (
	CmdHello  = "hello"
	CmdGoodby = "bye"
)

// Using a scanner to read the input
func Try1() {

	// This is the variable which is made for reading the inpput
	scanner := bufio.NewScanner(os.Stdin)

	// Creating variables for keeping track
	numLines := 0
	numCommands := 0

	// To keep the command running you need to have a infinite loop as practiced in hoemwork
	for scanner.Scan() { // this indicates that the data needs to be read

		/*
			strings.ToUpper is used to convert the string to upper case
		*/
		if strings.ToUpper(scanner.Text()) == "Q" {
			break // breaks from loop
		} else {
			text := strings.TrimSpace(scanner.Text()) // removes extra spaces
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("Cmd Response: Bastard")
			case CmdGoodby:
				numCommands += 1
				fmt.Println("Cmd Response: Fucker")
			}
			// Dont count blank lines
			if text != "" {
				numLines += 1
			}
		}
	}
	fmt.Println("Number of lines:", numLines)
	fmt.Println("Number of commands:", numCommands)

}
