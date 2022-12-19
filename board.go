package main

import (
	"math"
)

func FindDistanceBetweenCoords(c1 Coord, c2 Coord) int {
	var xDiff int
	var yDiff int

	xDiff = int(math.Abs(float64(c1.X - c2.X)))
	yDiff = int(math.Abs(float64(c1.Y - c2.Y)))

	return xDiff + yDiff
}

func FindPathBetweenCoords(c1 Coord, c2 Coord) Path {
	var path Path

	var xDiff int
	var yDiff int

	xDiff = int(math.Abs(float64(c1.X - c2.X)))
	yDiff = int(math.Abs(float64(c1.Y - c2.Y)))

	if c1.X < c2.X {
		for i := 1; i <= xDiff; i++ {
			path.Coords = append(path.Coords, Coord{X: c1.X + i, Y: c1.Y})
		}
	} else {
		for i := 1; i <= xDiff; i++ {
			path.Coords = append(path.Coords, Coord{X: c1.X - i, Y: c1.Y})
		}
	}

	if c1.Y < c2.Y {
		for i := 1; i <= yDiff; i++ {
			path.Coords = append(path.Coords, Coord{X: c2.X, Y: c1.Y + i})
		}
	} else {
		for i := 1; i <= yDiff; i++ {
			path.Coords = append(path.Coords, Coord{X: c2.X, Y: c1.Y - i})
		}
	}

	//log.Printf("Path: %v", path)

	return path
}

// Did not yet have the brain power to implement this.

func FindPathBetweenCoordsGoAroundSnakes(start Coord, end Coord, gs GameState) Path {
	var path Path

	return path
}

func CheckIfCoordIsSnake(coord Coord, gs GameState) bool {

	for i := 0; i < len(gs.Board.Snakes); i++ {
		// if the snake has NOT just eaten, it's tail is a safe spot to move
		if gs.Board.Snakes[i].Body[len(gs.Board.Snakes[i].Body)-1] != gs.Board.Snakes[i].Body[len(gs.Board.Snakes[i].Body)-2] {
			for j := 0; j < len(gs.Board.Snakes[i].Body)-1; j++ {
				if gs.Board.Snakes[i].Body[j] == coord {
					return true
				}
			}
		}
		// if the snake HAS just eaten, it's tail is NOT a safe spot to move
		if gs.Board.Snakes[i].Body[len(gs.Board.Snakes[i].Body)-1] == gs.Board.Snakes[i].Body[len(gs.Board.Snakes[i].Body)-2] {
			for j := 0; j < len(gs.Board.Snakes[i].Body); j++ {
				if gs.Board.Snakes[i].Body[j] == coord {
					return true
				}
			}
		}
	}
	return false
}

func CountFreeSpacesEachDirection(gs GameState, sm ScoredMoves) ScoredMoves {
	myHead := gs.You.Body[0]

	//log.Printf("Checking free spaces each direction")

	// Check right
	for i := 1; i <= 10; i++ {
		if myHead.X+i > gs.Board.Width {
			break
		}

		if CheckIfCoordIsSnake(Coord{X: myHead.X + i, Y: myHead.Y}, gs) {
			break
		}

		sm.Right.Score++
		sm.Right.Score += CountFreeNeighbours(gs, Coord{X: myHead.X + i, Y: myHead.Y})
	}

	// Check left
	for i := 1; i <= 10; i++ {
		if myHead.X-i < 0 {
			break
		}

		if CheckIfCoordIsSnake(Coord{X: myHead.X - i, Y: myHead.Y}, gs) {
			break
		}

		sm.Left.Score++
		sm.Left.Score += CountFreeNeighbours(gs, Coord{X: myHead.X - i, Y: myHead.Y})
	}

	// Check up
	for i := 1; i <= 10; i++ {
		if myHead.Y+i > gs.Board.Height {
			break
		}

		if CheckIfCoordIsSnake(Coord{X: myHead.X, Y: myHead.Y + i}, gs) {
			break
		}

		sm.Up.Score++
		sm.Up.Score += CountFreeNeighbours(gs, Coord{X: myHead.X, Y: myHead.Y + i})
	}

	// Check down
	for i := 1; i <= 10; i++ {
		if myHead.Y-i < 0 {
			break
		}

		if CheckIfCoordIsSnake(Coord{X: myHead.X, Y: myHead.Y - i}, gs) {
			break
		}

		sm.Down.Score++
		sm.Down.Score += CountFreeNeighbours(gs, Coord{X: myHead.X, Y: myHead.Y - i})
	}

	return sm
}

func GetNeighbourCoords(coord Coord) []Coord {
	return []Coord{
		{X: coord.X - 1, Y: coord.Y},
		{X: coord.X + 1, Y: coord.Y},
		{X: coord.X, Y: coord.Y - 1},
		{X: coord.X, Y: coord.Y + 1},
	}
}

func CountNeighboursThatDontHaveSnake(gs GameState, nb []Coord) int {
	var count int

	for _, coord := range nb {
		if !CheckIfCoordIsSnake(coord, gs) {
			count++
		}
	}

	return count
}

func CountAndReturnNeighboursThatDontHaveSnake(gs GameState, nb []Coord) ([]Coord, int) {
	var coords []Coord
	var count int

	for _, coord := range nb {
		if !CheckIfCoordIsSnake(coord, gs) {
			coords = append(coords, coord)
			count++
		}
	}

	return coords, count
}

func CountFreeNeighbours(gs GameState, coord Coord) int {
	var count int

	nb := GetNeighbourCoords(coord)

	for _, coord := range nb {
		if !CheckIfCoordIsSnake(coord, gs) {
			count++
		}
	}

	return count
}

func GetAllNeighbourCoords(coord Coord) []Coord {
	return []Coord{
		{X: coord.X - 1, Y: coord.Y},
		{X: coord.X + 1, Y: coord.Y},
		{X: coord.X, Y: coord.Y - 1},
		{X: coord.X, Y: coord.Y + 1},
		{X: coord.X - 1, Y: coord.Y - 1},
		{X: coord.X + 1, Y: coord.Y + 1},
		{X: coord.X - 1, Y: coord.Y + 1},
		{X: coord.X + 1, Y: coord.Y - 1},
	}
}

func CheckIfEnemyHeadIsInArea(eh Coord, area []Coord) bool {
	for _, coord := range area {
		if eh == coord {
			return true
		}
	}
	return false
}
