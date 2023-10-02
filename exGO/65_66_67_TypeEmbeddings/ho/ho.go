/*
Type _Embeddings - 
> Program is for warehouse , and determining automation for item types
*/

package ho

import (
	f "fmt"
	L "rape/lib"
)

func Meho() {
	f.Println("❤️")
	L.T1("❤️")
}

// Create constants  - Determines that conveyer belt on which items can move on 
const (
	Small = iota // iota automatically assigns a value for next variable
	Medium
	Large
)

// Second constants  - Determines ground and air 
const (
	Ground = iota
	Air
)

// Types for conveyer belt and shipping type 
type BeltSize int 
type Shipping int 

