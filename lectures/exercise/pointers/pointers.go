//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// - Security tags have two states: active (true) and inactive (false)
const (
	active   = true
	inactive = false
)

// * Create a structure to store items and their security tag state
type item struct {
	name string
	tag  bool
}

// * Create functions to activate and deactivate security tags using pointers
func activate(sectag *item) {
	sectag.tag = active
}

// * Create functions to activate and deactivate security tags using pointers
func deactivate(sectag *item) {
	sectag.tag = inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(shoppingList *[]item) {
	for i := 0; i < len(*shoppingList); i++ {
		deactivate(&(*shoppingList)[i])
	}
}

func printItems(items []item) {
	fmt.Println("--Item List--")
	var tags string
	for i := 0; i < len(items); i++ {
		if items[i].tag {
			tags = "Active"
		} else {
			tags = "Inactive"
		}
		fmt.Print(i+1, ".", items[i].name, ": ", tags, "   ")
	}
	fmt.Println()

}

func main() {
	// - Create at least 4 items, all with active security tags
	itemList := []item{{"Screw", inactive}, {"Nut", inactive}, {"Bolt", inactive}, {"Hex", inactive}}
	// - Store them in a slice or array
	printItems(itemList)

	for i := 0; i < len(itemList); i++ {
		item := &itemList[i]
		activate(item)
	}

	printItems(itemList)

	// - Deactivate any one security tag in the array/slice
	deactivate(&itemList[1])
	printItems(itemList)

	// - Call the checkout() function to deactivate all tags
	checkout(&itemList)
	printItems(itemList)

}

//  - Print out the array/slice after each change
