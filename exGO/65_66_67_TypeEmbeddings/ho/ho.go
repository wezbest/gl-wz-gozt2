/*
Type _Embeddings -
> Program is for warehouse , and determining automation for item types
*/

package ho

import (
	f "fmt"
)

func Meho() {

	// Main Function now implementing the stuff below
	mail := SpamMail{40000}
	Automate(&mail)

	// Toxic waste
	// waste := ToxicWaste{300}
	// Automate(&waste)
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

// Automation interface
type WarehouseAutomator interface {
	Conveyor
	Shipper
}

// TYpes of items that are available and whether or not they can be automated
type SpamMail struct {
	amount int
}

// String function what kind of object dealing with
func (s SpamMail) String() string {
	return "Spam Mail"
}

// Receiver function to implement interfaces
// Ship SpamMail by air
func (s *SpamMail) Ship() Shipping {
	return Air // Can be sent by air here
}

// Automation
func (s SpamMail) Convey() BeltSize {
	return Small // It should go on small belsize
}

// Automate function  using the warehouse interface
func Automate(item WarehouseAutomator) {
	f.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	f.Printf("Ship %v via %v\n", item, item.Ship())
}

// Something that cant be automated
type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}
