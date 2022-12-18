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
