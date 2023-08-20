package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Custom function to check if the mouse is over a rectangle
func IsMouseOverRectangle(rect rl.Rectangle) bool {
	mousePos := rl.GetMousePosition()
	return mousePos.X >= rect.X && mousePos.X <= rect.X+rect.Width && mousePos.Y >= rect.Y && mousePos.Y <= rect.Y+rect.Height
}

func main() {
	// Initialization
	var screenWidth, screenHeight int32 = 800, 600
	rl.InitWindow(screenWidth, screenHeight, "Coffee Tycoon")
	rl.SetTargetFPS(60)

	// Load the banner image and drinks spritesheet
	banner := rl.LoadTexture("assets/banner.png")
	drinkSprite := rl.LoadTexture("assets/drinks.png")

	// Initialize game state
	_ = drinksMade
	_ = ingredientsStock

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if IsMouseOverRectangle(rl.NewRectangle(0, float32(screenHeight-228), 64, 64)) {
				blackInstantCoffee.current_clicks++
				if blackInstantCoffee.current_clicks == blackInstantCoffee.make_time {
					drinksMade = append(drinksMade, blackInstantCoffee)
					blackInstantCoffee.current_clicks = 0

					// Decrement the ingredients stock for Black Instant Coffee
					for ingredient, amount := range blackInstantCoffee.ingredients {
						ingredientsStock[ingredient] -= amount
					}
				}
			} else if IsMouseOverRectangle(rl.NewRectangle(0, float32(screenHeight-164), 64, 64)) {
				whiteInstantCoffee.current_clicks++
				if whiteInstantCoffee.current_clicks == whiteInstantCoffee.make_time {
					drinksMade = append(drinksMade, whiteInstantCoffee)
					whiteInstantCoffee.current_clicks = 0

					// Decrement the ingredients stock for White Instant Coffee
					for ingredient, amount := range whiteInstantCoffee.ingredients {
						ingredientsStock[ingredient] -= amount
					}
				}
			}
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// Background
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Color{R: 245, G: 222, B: 179, A: 255})

		// Banner
		rl.DrawTexture(banner, 0, 0, rl.White)

		// Display drinks menu on the left side
		rl.DrawTextureRec(drinkSprite, rl.NewRectangle(0, 0, 64, 64), rl.NewVector2(0, float32(screenHeight-228)), rl.White)
		rl.DrawTextureRec(drinkSprite, rl.NewRectangle(0, 64, 64, 64), rl.NewVector2(0, float32(screenHeight-164)), rl.White)
		rl.DrawText(blackInstantCoffee.name, 70, int32(screenHeight-214), 20, rl.Black)
		rl.DrawText(strconv.Itoa(blackInstantCoffee.current_clicks)+"/"+strconv.Itoa(blackInstantCoffee.make_time), 70, int32(screenHeight-194), 20, rl.Black)
		rl.DrawText(whiteInstantCoffee.name, 70, int32(screenHeight-150), 20, rl.Black)
		rl.DrawText(strconv.Itoa(whiteInstantCoffee.current_clicks)+"/"+strconv.Itoa(whiteInstantCoffee.make_time), 70, int32(screenHeight-130), 20, rl.Black)
		rl.DrawText("Click to make coffee", int32(screenWidth/2)-100, int32(screenHeight-250), 20, rl.Black)

		rl.DrawText("Drinks Made: "+strconv.Itoa(len(drinksMade)), 10, 130, 20, rl.Black)
		for i, drink := range drinksMade {
			rl.DrawText(drink.name, 10, int32(150+(i*20)), 20, rl.Black)
		}

		// Display ingredients stock on the right side
		rl.DrawText("Ingredients Stock:", int32(screenWidth-200), 10, 20, rl.Black)
		rl.DrawText("Instant Coffee: "+strconv.Itoa(ingredientsStock[InstantCoffee{}]), int32(screenWidth-200), 40, 20, rl.Black)
		rl.DrawText("Milk: "+strconv.Itoa(ingredientsStock[Milk{}]), int32(screenWidth-200), 60, 20, rl.Black)
		rl.DrawText("Sugar: "+strconv.Itoa(ingredientsStock[Sugar{}]), int32(screenWidth-200), 80, 20, rl.Black)
		rl.DrawText("Ice Cubes: "+strconv.Itoa(ingredientsStock[IceCubes{}]), int32(screenWidth-200), 100, 20, rl.Black)
		rl.DrawText("Ice Cream: "+strconv.Itoa(ingredientsStock[IceCream{}]), int32(screenWidth-200), 120, 20, rl.Black)

		rl.EndDrawing()
	}

	// De-Initialization
	rl.UnloadTexture(banner)
	rl.UnloadTexture(drinkSprite)
	rl.CloseWindow()
}
