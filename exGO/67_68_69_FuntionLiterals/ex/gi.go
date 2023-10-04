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
	for _, line := range lines { // Original solution uses  c style loop 
		callback(line)
	}

	/*
	- This method uses the [i] , which represents an index 
	When asking cdeum - It said c style loop has moe control , obv since you can control 
	the steps via ( i++) and other conditions
	for lineIterator(lines []string, callback LineCallBack) {
		for i := 0; i < len(lines); i++ {
			callback(lines[i])
		}
	}
	*/
}

func mainFN() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	
	// These are the counters which are being used in lineFunc closure 
	letters := 0
	numbers := 0
	punct := 0
	spaces := 0

	// This is the closure which is being called in the lineIterator function
	lineFunc := func(line string) {

		// going through each line  and the function should run five times
		// because  we will iterate the lines slice with this closure 
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++ // Note this can also be written as letter +=1
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
