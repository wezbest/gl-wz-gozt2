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

// Iota patterns to get strings for BeltSize 
func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

// Iota patterns to get strings for Shipping Method 
func (b Shipping) String() string {
	return []string{"Ground", "Air"}[b]
}

// Creating interface for BeltSize 
type Conveyor interface {
	Convey() BeltSize 
}

// Create interface for choosing the type of shipping method 
type Shipper interface {
	Ship() Shipping
}

// 