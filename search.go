package main

import "log"

func BFS(start Coord, target Coord, gs GameState) Path {
	var path Path

	var queue []Coord
	queue = append(queue, start)

	neighbours := GetNeighbourCoords(start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if the node is the target
		if current == target {
			var targetPath Path
			log.Printf("targetPath: %v", targetPath)
			log.Printf("neighbours; %v", neighbours)
		}
	}
	return path
}
