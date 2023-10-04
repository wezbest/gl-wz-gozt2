/*
Excercise solution from q.txt
*/

package ex

import (
	"fmt"
	"unicode"
)

func Given() {
	mainFN()

}

// Given solution

// Create a type alias for a function
type LineCallBack func(line string)

// Actual function which is going to incorporate the type alias above
func lineIterator(lines []string, callback LineCallBack) {
	for _, line := range lines {
		callback(line)
	}
}

func mainFN() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	letters := 0
	numbers := 0
	punct := 0
	spaces := 0

	lineFunc := func(line string) {

		// going through each line  and the function should run five times
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++
			}
			if unicode.IsDigit(r) {
				numbers++
			}
			if unicode.IsPunct(r) {
				punct++
			}
			if unicode.IsSpace(r) {
				spaces++
			}
		}

	}
	lineIterator(lines, lineFunc)

	//Generate report
	fmt.Println(letters, "letters")
	fmt.Println(numbers, "Digits")
	fmt.Println(spaces, "Spaces")
	fmt.Println(punct, "punctuation")

}
