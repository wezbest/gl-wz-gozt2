/*
Given excercise

*/

package ex

import (
	"fmt"

	C "github.com/fatih/color"
)

// Main Function
func W1func() {

	// This is also given frrom q.txt
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	// Create dashboard
	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	// Create Colored text with fatih
	cre := C.New(C.FgBlue).SprintFunc()
	cye := C.New(C.FgYellow).SprintFunc()
	cbl := C.New(C.FgGreen).SprintFunc()

	//Printing out the dashboard
	fmt.Printf(cre("Average bandwidth is 	%v\n"), dash.AverageBandwidth())
	fmt.Printf(cbl("Average Cpu Temp is 	%v\n"), dash.AverageCpuTemp())
	fmt.Printf(cye("Average Memory Usage is %v\n"), dash.AverageMemoryUsage())

}

/*
Logic is here belopw
*/

type Bytes int
type Celcius float32

// -----------------------------------------------------------
/*
These are being embedded in line 22 
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}
And then they are being used in the receiver function in L35 - L37
*/
type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

// ------------------------------------------------------------

/*
This is the actual work which you are doing
*/

// Start by creating functions to calculate the averages

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}

// Average for Temp
func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}

// Average for memory
func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

// Creating the dashboard struct

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}
