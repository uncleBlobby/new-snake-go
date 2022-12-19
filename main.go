package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#2F2F2F", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	//fmt.Printf("State: %v", state)
	scoredMoves := ScoredMoves{}

	scoredMoves = AvoidNeck(state, scoredMoves)
	scoredMoves = AvoidWalls(state, scoredMoves)
	scoredMoves = AvoidSelf(state, scoredMoves)

	pathToCloseFood := FindPathBetweenCoords(state.You.Body[0], FindClosestFood(state))

	log.Printf("Path to closest food: %v", pathToCloseFood)

	scoredMoves = PreferPathTowardCloseFood(pathToCloseFood, state, scoredMoves)
	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	// opponents := state.Board.Snakes

	scoredMoves = AvoidSnakes(state, scoredMoves)

	//scoredMoves = CountFreeSpacesEachDirection(state, scoredMoves)

	scoredMoves = AvoidMovingInFrontOfLargerSnakes(state, scoredMoves)

	scoredMoves = CountOpenNodes(state, scoredMoves)

	// Choose the highest scored move to make
	nextMove := GetBestScoredMove(scoredMoves)

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	LogScoredMoves(scoredMoves)
	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)

	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
