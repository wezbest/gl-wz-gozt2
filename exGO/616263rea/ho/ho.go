/*
Readers 61 62 63 - Excercise
*/

package ho

import (
	"bufio"
	f "fmt"
	"io"
	"os"
	L "rape/lib"
	"strconv"
	"strings"
)

func Meho() {
	L.T1("Enter numbers and then exit with ctrl + d")
	mainFu()
}

// Create a reader

// This function will add the numbers hat have been asked as an input from the user

func mainFu() {

	// Creating a new reader - read from terminal ad numbers
	r := bufio.NewReader(os.Stdin)

	sum := 0

	// Infinite loop
	for {
		input, inputErr := r.ReadString(' ') // create slice of strings with rune
		n := strings.TrimSpace(input)        // remove extra spaces around
		if n == "" {                         // Check if only spaces was input
			continue // skip spaces
		}

		//Convert string to number
		num, convErr := strconv.Atoi(n) // convert string to int
		if convErr != nil {
			f.Println(convErr)
		} else { // if no error add sum
			sum += num
		}

		// check for no more strings
		if inputErr == io.EOF {
			break
		}

		// Check if actual error occurred
		if inputErr != nil {
			f.Println("Error reading stdin: Bastard what u do fucker ? : ", inputErr)
		}

	}

	// Now just print the sum out
	f.Printf("sum: %v\n", sum)

}
