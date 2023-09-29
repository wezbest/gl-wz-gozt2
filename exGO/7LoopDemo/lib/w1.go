/*
22 Loop Demo
*/

package lib

import (
	f "fmt"
)

func Fu1() {
	sum := 0
	f.Println("Sum Is Suck", sum)

	// Here 'i' keeps getting incremented until it is <= 10 , then exits 
	for i := 1; i <= 10; i++ {
		sum += i
		f.Println("Sum Is Suck", sum)
	}

	/*
	Using the sum variable above , we are just removing 5 util it is less than 10 
	This works because from the L18 , the value of sum is 55 and is starts  out from there 
	*/
	for sum > 10 {
		sum -= 5
		f.Println("Decrement", sum)
	}
}
