package main

import "log"

func LogScoredMoves(sm ScoredMoves) {
	//log.Printf("ScoredMoves: %v", sm)

	log.Printf("ScoredMoves.Left: %v", sm.Left)
	log.Printf("ScoredMoves.Up: %v", sm.Up)
	log.Printf("ScoredMoves.Right: %v", sm.Right)
	log.Printf("ScoredMoves.Down: %v", sm.Down)
}
