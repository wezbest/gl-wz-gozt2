/*
Given solution
*/

package lib

import (
	f "fmt"
)

// This type was already given

type Part string

// Print contents

func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i] // This is making copy to prevent cooncurrency issues later on
		f.Println(part)
	}
}

// Main solution
func GiMain() {

	// Making an empty slice with required elements
	Texc("Making empty slice and printing contents")
	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "bolt"
	assemblyLine[2] = "Sheets"

	f.Println("3 Parts:")
	showLine(assemblyLine)

	// Adding two new parts
	Texc("Adding two new parts")
	assemblyLine = append(assemblyLine, "washet", "wheel")
	f.Println("\n Added two new parts")
	showLine(assemblyLine)

	// Slice assembly line contains two new parts
	Texc("Slice assembly line contains two new parts")
	assemblyLine = assemblyLine[3:]
	f.Println("\n Slice assembly line contains two new parts")
	showLine(assemblyLine)

}
