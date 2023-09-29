/*
tests for givete.go , which is the given solution for the receiver function exercise
*/

package givte

import (
	"testing"
)

// Helper funciton

func newPlayer() Player2 {
	return Player2{
		name:      "AssLicker",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

}

// Testing the health

func TestHealth(t *testing.T) {
	player := newPlayer()
	// Adding health more that max
	player.addHealth(999)
	if player.health > player.maxHealth {
		t.Fatalf("Health wen over limit: %v, want %v", player.health, player.maxHealth)
	}

	// Testing the damage
	// Which is max health +1
	player.applyDamage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("Health: %v. Minimim: 0", player.health)
	}

	// Checking overflow 
	if player.health > player.maxHealth {
		t.Fatalf("health: %v. Max: %v", player.health, player.maxHealth)
	}
}

// Same like above test but were are checking for the energy , overflow and wrap around
func TestEnergy(t *testing.T) {
	player := newPlayer()
	// Adding health more that max
	player.addEnergy(999)
	if player.energy > player.maxEnergy {
		t.Fatalf("Energy wen over limit: %v, want %v", player.energy, player.maxEnergy)
	}

	// Testing the damage
	// Which is max health +1
	player.consumeEnergy(player.maxEnergy + 1)
	if player.health < 0 {
		t.Fatalf("Energy: %v. Minimim: 0", player.energy)
	}

	// Checking overflow 
	if player.energy > player.maxEnergy {
		t.Fatalf("Energy: %v. Max: %v", player.energy, player.maxEnergy)
	}
}