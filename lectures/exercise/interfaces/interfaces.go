//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

type Sorter interface {
	sort() Lift
}

type Truck string
type Car string
type Motorcycle string
type Lift int

func (m Motorcycle) sort() Lift {
	return 0
}
func (m Car) sort() Lift {
	return 1
}
func (m Truck) sort() Lift {
	return 2
}

type Vehicle struct {
	bikes  []Motorcycle
	cars   []Car
	trucks []Truck
}

var (
	garage Vehicle
)

func init() {
	garage = Vehicle{
		bikes:  []Motorcycle{"RoyalEnfield", "Bullet", "Cobalt"},
		cars:   []Car{"Baleno", "Pajero", "Hyundai"},
		trucks: []Truck{"Silverado", "Lumberjack", "Dominator"},
	}

}

// These functions have been copied from JL's solution. Need to understand why we need to call the String() functions below.
// Without these functions we can only type the value type of motorcycle as main.motorcycle when we only want motercycle.
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func sendInVehicle(m Sorter) {
	switch m.sort() {

	case 0:
		fmt.Printf("Send %v to the small lift.\n", m)
	case 1:
		fmt.Printf("Send %v to the standard lift.\n", m)
	case 2:
		fmt.Printf("Send %v to the large lift.\n", m)
	}
}

func main() {
	m := garage.bikes[1]
	sendInVehicle(m)
}
