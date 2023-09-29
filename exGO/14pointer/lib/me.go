/*
My attempt at a solution
Requirements
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

*/

package lib

import (
	"fmt"
)

func M1f() {
	// Creating two seperate arrays
	item := []string{"Item 1", "Item 2", "Item 3", "Item 4"}
	tag := []bool{true, true, true, true}
	// Deactivating the first item
	tag[0] = false
	fmt.Println(item, tag)
	m1chk(tag)
	fmt.Println(item, tag)
}

func m1chk(checkout []bool) {
	for i := 0; i < len(checkout); i++ {
		checkout[i] = false 
	}
}
