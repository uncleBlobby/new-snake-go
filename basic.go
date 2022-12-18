package main

import "log"

func AvoidNeck(gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]
	myNeck := gs.You.Body[1]

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		sm.Left.Score = -1000

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		sm.Right.Score = -1000

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		sm.Down.Score = -1000

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		sm.Up.Score = -1000
	}

	return sm
}

func AvoidWalls(gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]

	if myHead.X == 0 { // Head is on left wall, don't move left
		sm.Left.Score = -1000
	}

	if myHead.X == gs.Board.Width-1 { // Head is on right wall, don't move right
		sm.Right.Score = -1000
	}

	if myHead.Y == 0 { // Head is on top wall, don't move up
		sm.Down.Score = -1000
	}

	if myHead.Y == gs.Board.Height-1 { // Head is on bottom wall, don't move down
		sm.Up.Score = -1000
	}

	return sm
}

func AvoidSelf(gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]

	for _, bodyPart := range gs.You.Body[1:] {

		if myHead.X+1 == bodyPart.X && myHead.Y == bodyPart.Y {
			log.Printf("Avoiding self (RIGHT): %v", bodyPart)
			sm.Right.Score = -1000
		}

		if myHead.X-1 == bodyPart.X && myHead.Y == bodyPart.Y {
			log.Printf("Avoiding self (LEFT): %v", bodyPart)
			sm.Left.Score = -1000
		}

		if myHead.X == bodyPart.X && myHead.Y+1 == bodyPart.Y {
			log.Printf("Avoiding self (UP): %v", bodyPart)
			sm.Up.Score = -1000
		}

		if myHead.X == bodyPart.X && myHead.Y-1 == bodyPart.Y {
			log.Printf("Avoiding self (DOWN): %v", bodyPart)
			sm.Down.Score = -1000
		}
	}

	return sm
}

func AvoidSnakes(gs GameState, sm ScoredMoves) ScoredMoves {
	myHead := gs.You.Body[0]

	for _, snake := range gs.Board.Snakes {
		for _, bodyPart := range snake.Body {

			// if snake has just eaten, avoid their tail
			if snake.Body[len(snake.Body)-1] == snake.Body[len(snake.Body)-2] {
				//log.Printf("%v has just eaten, avoiding tail. ", snake.Name)
				if myHead.X+1 == bodyPart.X && myHead.Y == bodyPart.Y {
					log.Printf("Avoiding snake (RIGHT): %v", bodyPart)
					sm.Right.Score = -1000
				}

				if myHead.X-1 == bodyPart.X && myHead.Y == bodyPart.Y {
					log.Printf("Avoiding snake (LEFT): %v", bodyPart)
					sm.Left.Score = -1000
				}

				if myHead.X == bodyPart.X && myHead.Y+1 == bodyPart.Y {
					log.Printf("Avoiding snake (UP): %v", bodyPart)
					sm.Up.Score = -1000
				}

				if myHead.X == bodyPart.X && myHead.Y-1 == bodyPart.Y {
					log.Printf("Avoiding snake (DOWN): %v", bodyPart)
					sm.Down.Score = -1000
				}
			}

			// if snake has not just eaten, their tail is safe
			if snake.Body[len(snake.Body)-1] != snake.Body[len(snake.Body)-2] {
				if bodyPart != snake.Body[len(snake.Body)-1] {
					if myHead.X+1 == bodyPart.X && myHead.Y == bodyPart.Y {
						log.Printf("Avoiding snake (RIGHT): %v", bodyPart)
						sm.Right.Score = -1000
					}

					if myHead.X-1 == bodyPart.X && myHead.Y == bodyPart.Y {
						log.Printf("Avoiding snake (LEFT): %v", bodyPart)
						sm.Left.Score = -1000
					}

					if myHead.X == bodyPart.X && myHead.Y+1 == bodyPart.Y {
						log.Printf("Avoiding snake (UP): %v", bodyPart)
						sm.Up.Score = -1000
					}

					if myHead.X == bodyPart.X && myHead.Y-1 == bodyPart.Y {
						log.Printf("Avoiding snake (DOWN): %v", bodyPart)
						sm.Down.Score = -1000
					}
				}

			}

		}
	}

	return sm
}

func AvoidMovingInFrontOfLargerSnakes(gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]

	for _, snake := range gs.Board.Snakes {
		if len(snake.Body) < len(gs.You.Body) {
			continue
		}

		if snake.ID == gs.You.ID {
			continue
		}

		enemyHead := snake.Body[0]

		myLeft := Coord{X: myHead.X - 1, Y: myHead.Y}
		myRight := Coord{X: myHead.X + 1, Y: myHead.Y}
		myUp := Coord{X: myHead.X, Y: myHead.Y + 1}
		myDown := Coord{X: myHead.X, Y: myHead.Y - 1}

		// left

		if CheckIfEnemyHeadIsInArea(enemyHead, GetAllNeighbourCoords(myLeft)) {
			log.Printf("Avoiding enemy left (LARGEENEMYHEADINAREA): %v", enemyHead)
			sm.Left.Score = -100
		}

		// right

		if CheckIfEnemyHeadIsInArea(enemyHead, GetAllNeighbourCoords(myRight)) {
			log.Printf("Avoiding enemy right (LARGEENEMYHEADINAREA): %v", enemyHead)
			sm.Right.Score = -100
		}

		// up

		if CheckIfEnemyHeadIsInArea(enemyHead, GetAllNeighbourCoords(myUp)) {
			log.Printf("Avoiding enemy up (LARGEENEMYHEADINAREA): %v", enemyHead)
			sm.Up.Score = -100
		}

		// down

		if CheckIfEnemyHeadIsInArea(enemyHead, GetAllNeighbourCoords(myDown)) {
			log.Printf("Avoiding enemy down (LARGEENEMYHEADINAREA): %v", enemyHead)
			sm.Down.Score = -100
		}

	}

	return sm
}

func PreferPathTowardCloseFood(path Path, gs GameState, sm ScoredMoves) ScoredMoves {

	myHead := gs.You.Body[0]
	distanceToFood := len(path.Coords)

	healthRemaining := gs.You.Health

	if len(path.Coords) == 0 {
		return sm
	}

	if healthRemaining <= distanceToFood+10 {

		// Prefer the move that will get us closer to the closest food
		if path.Coords[0].X > myHead.X {

			sm.Right.Score += 100
		}

		if path.Coords[0].X < myHead.X {
			sm.Left.Score += 100
		}

		if path.Coords[0].Y > myHead.Y {
			sm.Up.Score += 100
		}

		if path.Coords[0].Y < myHead.Y {
			sm.Down.Score += 100
		}
	}
	return sm
}

func GetBestScoredMove(sm ScoredMoves) string {
	bestMove := "up"
	bestScore := sm.Up.Score

	if sm.Down.Score > bestScore {
		bestScore = sm.Down.Score
		bestMove = "down"
	}

	if sm.Left.Score > bestScore {
		bestScore = sm.Left.Score
		bestMove = "left"
	}

	if sm.Right.Score > bestScore {
		bestScore = sm.Right.Score
		bestMove = "right"
	}

	log.Printf("Best move: %s", bestMove)
	log.Printf("Best score: %d", bestScore)
	return bestMove
}
