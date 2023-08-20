package main

// Drink struct
type Drink struct {
	name           string
	ingredients    map[interface{}]int
	make_time      int
	current_clicks int
}

// Drinks Menu

var blackInstantCoffee = Drink{
	name: "Black Instant Coffee",
	ingredients: map[interface{}]int{
		InstantCoffee{}: 30,
	},
	make_time:      20,
	current_clicks: 0,
}

var whiteInstantCoffee = Drink{
	name: "White Instant Coffee",
	ingredients: map[interface{}]int{
		InstantCoffee{}: 30,
		Milk{}:          100,
	},
	make_time:      30,
	current_clicks: 0,
}
