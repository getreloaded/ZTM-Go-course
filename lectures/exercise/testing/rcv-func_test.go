//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package testing

import "testing"

func testPlayer() Player {
	return Player{Name: "KB", Health: 100, Energy: 50, MaxHealth: 100, MaxEnergy: 100}
}

func TestDeltaHealth(t *testing.T) {

	newTest := []struct {
		delta     int
		testvalue int
	}{
		{200, 100},
		{-200, 0},
		{50, 50},
		{-25, 25},
	}
	newPlayer := testPlayer()
	for i := 0; i < len(newTest); i++ {
		deltaTest := newTest[i].delta
		newPlayer.DeltaHealth(deltaTest)
		if newPlayer.Health != newTest[i].testvalue {
			t.Errorf("Health has crossed the limits. Expected %v, got %v", newTest[i].testvalue, newPlayer.Health)
		}
	}

}
