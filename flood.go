package main

import "log"

func CountOpenNodes(gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]

	clearanceLeft := 0
	clearanceRight := 0
	clearanceUp := 0
	clearanceDown := 0

	for i := myHead.X - 1; i >= 0; i-- {
		if !CheckIfCoordIsSnake(Coord{X: i, Y: myHead.Y}, gs) {
			clearanceLeft++
			for j := myHead.Y + 1; j < gs.Board.Height; j++ {
				if !CheckIfCoordIsSnake(Coord{X: i, Y: j}, gs) {
					clearanceLeft++
				} else {
					break
				}
			}
			for j := myHead.Y - 1; j >= 0; j-- {
				if !CheckIfCoordIsSnake(Coord{X: i, Y: j}, gs) {
					clearanceLeft++
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	for i := myHead.X + 1; i < gs.Board.Width; i++ {
		if !CheckIfCoordIsSnake(Coord{X: i, Y: myHead.Y}, gs) {
			clearanceRight++
			for j := myHead.Y + 1; j < gs.Board.Height; j++ {
				if !CheckIfCoordIsSnake(Coord{X: i, Y: j}, gs) {
					clearanceRight++
				} else {
					break
				}
			}
			for j := myHead.Y - 1; j >= 0; j-- {
				if !CheckIfCoordIsSnake(Coord{X: i, Y: j}, gs) {
					clearanceRight++
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	for i := myHead.Y - 1; i >= 0; i-- {
		if !CheckIfCoordIsSnake(Coord{X: myHead.X, Y: i}, gs) {
			clearanceDown++
			for j := myHead.X + 1; j < gs.Board.Width; j++ {
				if !CheckIfCoordIsSnake(Coord{X: j, Y: i}, gs) {
					clearanceDown++
				} else {
					break
				}
			}
			for j := myHead.X - 1; j >= 0; j-- {
				if !CheckIfCoordIsSnake(Coord{X: j, Y: i}, gs) {
					clearanceDown++
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	for i := myHead.Y + 1; i < gs.Board.Height; i++ {
		if !CheckIfCoordIsSnake(Coord{X: myHead.X, Y: i}, gs) {
			clearanceUp++
			for j := myHead.X + 1; j < gs.Board.Width; j++ {
				if !CheckIfCoordIsSnake(Coord{X: j, Y: i}, gs) {
					clearanceUp++
				} else {
					break
				}
			}
			for j := myHead.X - 1; j >= 0; j-- {
				if !CheckIfCoordIsSnake(Coord{X: j, Y: i}, gs) {
					clearanceUp++
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	log.Printf("Clearance Left: %d\n", clearanceLeft)
	log.Printf("Clearance Right: %d\n", clearanceRight)
	log.Printf("Clearance Up: %d\n", clearanceUp)
	log.Printf("Clearance Down: %d\n", clearanceDown)

	sm.Left.Score += clearanceLeft
	sm.Right.Score += clearanceRight
	sm.Up.Score += clearanceUp
	sm.Down.Score += clearanceDown

	return sm
}
