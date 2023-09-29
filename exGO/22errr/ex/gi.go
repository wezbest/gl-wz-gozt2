/*
Given solution
*/

package ex

import (
	f "fmt"
	"strconv"
	"strings"
)

func Given() {
	f.Println("❤️")
}

// Create a time structure

type Time struct {
	hrs, mins, secs int
}

//Custom error type , instead of using error.new from the error package

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	// The t.msg is a pointer to the msg field fro the struct being made above , note how pointer is being used
	return f.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	// Parse the input string to get the time components

	components := strings.Split(input, ":")

	// This case is to check if the time input is in the correct format
	if len(components) != 3 {
		// Check if the right number of components
		return Time{}, &TimeParseError{"invalid time", input}
	} else {
		//Parsing hour into integer
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{f.Sprintf("Error parsing hour: %v", err), input}
		}

		//Parsing Minute into integer
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{f.Sprintf("Error parsing minute: %v", err), input}
		}

		//Parsing Second into integer
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{f.Sprintf("Error parsing seconds: %v", err), input}
		}

		//check numbers that are being input
		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour of out range: 0 <= hour <= 23", f.Sprintf("%v", hour)}
		}
		//check minutes
		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{"minute of out range: 0 <= minute <= 59", f.Sprintf("%v", minute)}
		}
		//check seconds
		if second > 59 || second < 0 {
			return Time{}, &TimeParseError{"second of out range: 0 <= second <= 59", f.Sprintf("%v", second)}
		}

		return Time{hour, minute, second}, nil
	}
}
