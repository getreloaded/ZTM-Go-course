//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:

package testing

import "fmt"

// * Implement a player having the following statistics:
//   - Health, Max Health
//   - Energy, Max Energy
//   - Name
type Player struct {
	Health, MaxHealth int
	Energy, MaxEnergy int
	Name              string
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.

func (player *Player) DeltaHealth(changeInHealth int) {
	player.Health += changeInHealth
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
	if player.Health < 0 {
		player.Health = 0
	}

	// - Print out the statistic change within each function
	fmt.Println("Change health:", changeInHealth)
}

func (player *Player) DeltaEnergy(changeInEnergy int) {
	player.Energy += changeInEnergy
	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}
	if player.Energy < 0 {
		player.Energy = 0
	}

	// - Print out the statistic change within each function
	fmt.Println("Change Energy:", changeInEnergy)
}

// - Execute each function at least once
func main() {

	newPlayer := Player{Name: "KB", Health: 100, Energy: 50, MaxHealth: 100, MaxEnergy: 100}
	fmt.Println(newPlayer.Name, " Stats --- Health:", newPlayer.Health, "and Energy:", newPlayer.Energy)

	newPlayer.DeltaHealth(-20)
	newPlayer.DeltaEnergy(10)
	fmt.Println(newPlayer.Name, " Stats --- Health:", newPlayer.Health, "and Energy:", newPlayer.Energy)

	newPlayer.DeltaHealth(-100)
	newPlayer.DeltaEnergy(200)
	fmt.Println(newPlayer.Name, " Stats --- Health:", newPlayer.Health, "and Energy:", newPlayer.Energy)

}
