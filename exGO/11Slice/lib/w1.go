/*
Actual work will go here - Print out a route for multiple destinations 
*/

package lib 

import (
	f "fmt"
)

// Defining slice in the following function 
func W1f() {
	Texc("Original slice which was assigned")
	route  := []string{"New York", "Paris", "London", "Tokyo"}
	WIp("Route1", route)
	
	// Npw to append new element into the slice 
	Texc("Added Berlin which has been appended into the slice")
	route = append(route, "Berlin")
	WIp("Route2", route)
	
	// Now we will visit some places - This is using the slicing syntax 
	Texc("Using selecting slice element below")
	f.Println("")
	f.Println(route[0], "visited") // This will bring NewYork which in index 0 
	f.Println(route[1], "visited") // This will bring Paris which in index Paris 
	
	// eslice existing slice
	Texc("Reslice the slice , and printing from index2")
	route = route[2:] // slice syntax include from index 2 which is tokyo berlin
	WIp("Remaining locations", route) 
}

// Helper function to print out the slice 
func WIp(title string , slice []string) {
	f.Println()
	f.Println("---", title,"---")
	
	// Writing a loop iterate and print out the slice
	for i:=0; i<len(slice); i++ {
		element := slice[i] // This is making copy to prevent cooncurrency issues later on 
		f.Println(element)
	}
}