/*
Given solution
- /sr-dice project
*/

package lib

import (
	f "fmt"
	r "math/rand"
	t "time"
)

func Gi1() {
	// Seed Random number
	r.Seed(t.Now().UnixNano())

	dice, sides := 2, 12
	rolls := 1

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := Groll(sides)
			sum += rolled
			f.Println("Roll #", r, ", Die #", d, ": ", rolled)
		}
		f.Println("Total #", r, ": ", sum)
		switch sum := sum; {
		case sum == 2:
			f.Println("Snake Eyes")
		case sum == 7:
			f.Println("Lucky 7")
		case sum%2 == 0:
			f.Println("Even")
	    case sum%2 != 1:
			f.Println("Odd")
		}
	}
}

func Gi2() {
	// Seed Random number
	r.Seed(t.Now().UnixNano())

	dice, sides := 2, 12
	rolls := 5

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := Groll(sides)
			sum += rolled
			f.Println("Roll #", r, ", Die #", d, ": ", rolled)
		}
		f.Println("Total #", r, ": ", sum)
		switch sum := sum; {
		case sum == 2:
			f.Println("Snake Eyes")
		case sum == 7:
			f.Println("Lucky 7")
		case sum%2 == 0:
			f.Println("Even")
	    case sum%2 != 1:
			f.Println("Odd")
		}
	}
}

func Groll(sides int) int {
	return r.Intn(sides) + 1
}
