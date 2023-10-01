/*
Ranges learning 
- This is basically an easy way of iterating over a slice which was also mentioned in the 
original GoLangAssLickingTutorial which you did earlier
*/

package lib 

import (
	f "fmt"
)

type ranga string

func Wo1() {
	slice := []ranga{"Hello", "Ass", "Liqer"}

	for i, element := range slice {
		f.Println(i, element, ":")
		
		// Using %q will print out the runes for each letter - runes are 
		for _, ch := range element {
			f.Printf("  %q\n", ch)
		}
	}
}