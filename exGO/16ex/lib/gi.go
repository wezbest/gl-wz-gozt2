/*
Given solution
*/

package lib

import (
	f "fmt"

	C "github.com/fatih/color"
)

// This is the main function which will have all the other functions below
func Given() {

	player := Player2{
		name:      "AssLicker",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	// Printing Original before additions
	ccy := C.New(C.FgCyan, C.Bold).SprintFunc()
	f.Println(ccy("Original Before: "), ccy(player))

	// Making Changes to the stats here

	player.applyDamage(99)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(10)

	// Overflow behavior test
	player.consumeEnergy(9999)

	// Printing the original again after edits
	crey := C.New(C.FgYellow, C.Bold).SprintFunc()
	f.Println("\n", crey("Original After Additions:"), crey(player), "\n")
}

/*
The function below are going to be called above
*/

// Creating player structure - This is actually made the requirements

type Player2 struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

// Adding health function
func (player *Player2) addHealth(amount uint) {

	player.health += amount
	// This condition is if the health is more that moaxhealth then set to maxHealth
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	f.Println(player.name, "Add", amount, "health ->", player.health)
}

// Adding Energy  function
func (player *Player2) addEnergy(amount uint) {

	player.energy += amount
	// This condition is if the health is more that moaxhealth then set to maxHealth
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	f.Println(player.name, "Add", amount, "energy ->", player.energy)
}

// Adding damage - which means subtracting the health
func (player *Player2) applyDamage(amount uint) {

	// If the amount of Damage being added is more that the health, then make it 0
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	f.Println(player.name, "Damage", amount, "Damage ->", player.health)
}

// Consuming Energy - which means subtracting the energy
func (player *Player2) consumeEnergy(amount uint) {

	// If the amount of Damage being added is more that the health, then make it 0
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	f.Println(player.name, "Consume ", amount, "Consuming Energy ->", player.energy)
}
