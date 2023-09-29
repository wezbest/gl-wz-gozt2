/*
My attemtp at excercise
Problem


*/

package lib

import (
	f "fmt"

	C "github.com/fatih/color"
)

func Me() {
	// Define the player first
	player := Player{
		Name:      "John",  
		Health:    50,
		Energy:    50,
		MaxHealth: 100,
		MaxEnergy: 100,
	}

	//Original
	cre := C.New(C.FgRed).SprintFunc()
	f.Println(cre("Original :"), cre(player), "\n")

	//Run Modification
	player.ModifyHealth(10)
	cbl := C.New(C.FgBlue).SprintFunc()
	f.Println(cbl("Modifying Health +10 (2nd Value): "), cbl(player))

	// Modify Energy
	player.ModifyEnergy(10)
	cgr := C.New(C.FgGreen).SprintFunc()
	f.Println(cgr("Modifying Energy +10 (3rd Value) : "), cgr(player))

	// Modifying Max MaxHealth
	player.ModifyMaxHealth(100)
	cmg := C.New(C.FgMagenta, C.Bold).SprintFunc()
	f.Println(cmg("Modifying Max Health +100 (4th Value): "), cmg(player))

	// Modifying Max Energy
	player.ModifyMaxEnergy(100)
	ccy := C.New(C.FgCyan, C.Bold).SprintFunc()
	f.Println(ccy("Modifying Max Health +100 (5th Value): "), ccy(player))

	// Printing Original Again
	//Original
	crey := C.New(C.FgYellow, C.Bold).SprintFunc()
	f.Println("\n", crey("Original After Additions:"), crey(player), "\n")
}

// Creating the variables

type Player struct {
	Name      string
	Health    int
	Energy    int
	MaxHealth int
	MaxEnergy int
}

//Color Definitions

// Create receiver functions modify statistics

func (p *Player) ModifyHealth(health int) {
	p.Health += health
}

func (p *Player) ModifyEnergy(energy int) {
	p.Energy += energy
}

func (p *Player) ModifyMaxEnergy(MaxEnergy int) {
	p.MaxEnergy += MaxEnergy
}
func (p *Player) ModifyMaxHealth(MaxHealth int) {
	p.MaxHealth += MaxHealth
}
