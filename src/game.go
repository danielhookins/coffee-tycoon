package main

// Initial game state
var drinksMade []Drink

var ingredientsStock = map[interface{}]int{
	InstantCoffee{}: 1000,
	Milk{}:          1000,
	Sugar{}:         1000,
	IceCubes{}:      0,
	IceCream{}:      0,
}
