package main

import (
	"log"
)

func FindClosestFood(gs GameState) Coord {
	myHead := gs.You.Body[0]

	var closestFood Coord = gs.Board.Food[0]

	for _, food := range gs.Board.Food {
		if FindDistanceBetweenCoords(myHead, food) < FindDistanceBetweenCoords(myHead, closestFood) {
			closestFood = food
		}
	}

	log.Printf("Closest food: %v", closestFood)
	return closestFood
}

// Not sure if the below function works properly, haven't tried it yet.

func FindClosestFoodThatNoOtherSnakeIsCloseTo(gs GameState) (Coord, bool) {
	myHead := gs.You.Body[0]

	thisFoodIsSafe := false

	var closestFood Coord = FindClosestFood(gs)

	for _, food := range gs.Board.Food {

		if FindDistanceBetweenCoords(myHead, food) < FindDistanceBetweenCoords(myHead, closestFood) {
			for _, snake := range gs.Board.Snakes {
				if FindDistanceBetweenCoords(snake.Body[0], food) < FindDistanceBetweenCoords(myHead, food) {
					thisFoodIsSafe = false
					break

				}
				if FindDistanceBetweenCoords(snake.Body[0], food) > FindDistanceBetweenCoords(myHead, food) {
					closestFood = food
					thisFoodIsSafe = true
				}
			}

		}
	}

	log.Printf("Closest food: %v is safe? %v", closestFood, thisFoodIsSafe)
	return closestFood, thisFoodIsSafe
}
