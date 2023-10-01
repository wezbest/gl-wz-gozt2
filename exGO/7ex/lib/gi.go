/*
Given solution
*/

package lib

import (
	f "fmt"
)

func Gi() {

	// Print integers 1 to 50, except:
	for i := 1; i <= 50; i++ {

		//  - Print "Fizz" if the integer is divisible by 3
		//  - Print "Buzz" if the integer is divisible by 5
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5

		// First check if it is divisible by 3 as compared to yours which was all in 1 block
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		// For  numbers it is important to have it in the correct order
		if divisibleBy3 && divisibleBy5 {
			T1("FizzBuzz - div3div5")
		} else if divisibleBy3 {
			T2("Fizz - div3")
		} else if divisibleBy5 {
			T3("Buzz - div5")
		} else {
			f.Println(i)
		}

	}

}
