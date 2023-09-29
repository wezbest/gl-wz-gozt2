/*
39 Pointers in GO
*/

package lib

import (
	f "fmt"
)

func W1f() {
	counter := Counter{}
	hello := "Hello"
	world := "World"
	f.Println(hello, world)

	// Here the replacement is  being done
	replace (&hello, "Goodbye", &counter)
	f.Println(hello, world)

	// Creating a salice now and implement pointers 
	phrase := []string{hello, world}
	f.Println(phrase)
 
	// replace phrase 
	replace(&phrase[0], "Goodbye", &counter)
	f.Println(phrase)
}

type Counter struct {
	hits int
}

// Function that accepts a counter pointer variable which is a struct
func increment(counter *Counter) {
	counter.hits += 1 // here pointer assignment not required since structs automatically to do that 
	f.Println("Counter", counter)
}

/*
1. Replaced old string with someting else 
2. Incremented the counter
*/

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}