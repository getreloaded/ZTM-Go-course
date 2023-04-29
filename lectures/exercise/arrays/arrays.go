//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

const listlen = 4

type Product struct {
	name  string
	price float64
}

func printInfo(shoppingList [listlen]Product) {
	totalItems := 0
	totalCost := 0.0
	for i := 0; i < listlen; i++ {
		j := i
		if shoppingList[j].name != "" {
			totalItems++
			totalCost += shoppingList[j].price
		}
	}

	fmt.Printf("The last item on the list is %v\n", shoppingList[totalItems-1].name)
	fmt.Printf("The total number of items is %v\n", totalItems)
	fmt.Printf("The total cost of the items is %v\n", totalCost)

}

func main() {

	shoppingList := [listlen]Product{}

	shoppingList[0] = Product{
		name:  "Apple",
		price: 100,
	}

	shoppingList[1] = Product{
		name:  "Banana",
		price: 200,
	}

	shoppingList[2] = Product{
		name:  "Cherry",
		price: 300,
	}

	printInfo(shoppingList)

	shoppingList[3] = Product{
		name:  "Orange",
		price: 400,
	}
	printInfo(shoppingList)

}
