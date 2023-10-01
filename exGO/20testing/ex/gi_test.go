/*
Given solution tests
*/

package ex

import (
	f "fmt"
	"testing"

	C "github.com/fatih/color"
)

// Testing add energy from L67

func TestAddEnergy(t *testing.T) {
	player := Player2{
		name:      "AssLicker",
		health:    100,
		maxHealth: 100,
		energy:    0,
		maxEnergy: 500,
	}

	// First lets print out the original
	ccy := C.New(C.FgCyan, C.Bold).SprintFunc()
	cred := C.New(C.FgRed, C.Bold).SprintFunc()
	f.Println(ccy("Original Before: "), ccy(player.energy))

	// Now lets add and do testing
	player.addEnergy(100)
	if player.energy != 100 {
		t.Errorf("Failed to add energy")
	} else {
		f.Println(cred("Addition Pass"), cred(player.energy))
	}

}
