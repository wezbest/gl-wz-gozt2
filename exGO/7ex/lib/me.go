/*
My solution for the loops

/* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
*/

package lib 

import (
	f "fmt"
)

func M1f() {
	for i := 1; i <= 50; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			f.Println("FizzBuzz")
		} else if i % 3 == 0 {
			f.Println("Fizz")
		} else if i % 5 == 0 {
			f.Println("Buzz")
		} else {
			f.Println(i)
		}
	}
}
