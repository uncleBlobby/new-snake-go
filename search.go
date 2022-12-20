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

func BreadthFirstSearch(start Coord, target Coord, gs GameState) Path {
	var path Path

	var startNode SearchNode = SearchNode{Coord: start, Parent: nil}

	var queue []SearchNode
	queue = append(queue, startNode)

	var visited []SearchNode
	visited = append(visited, startNode)

	for len(queue) > 0 {
		//log.Printf("queue: %v", queue)
		log.Printf("made it here")
		log.Printf("queue length: %v", len(queue))
		log.Printf("have visited: %v", visited)
		current := queue[0]
		queue = queue[1:]

		// Check if the node is the target
		if current.Coord == target {
			var targetPath Path
			for current.Parent != nil {
				targetPath.Coords = append(targetPath.Coords, current.Coord)
				current = *current.Parent
			}
			//reverse the targetPath and store it in path
			for i := len(targetPath.Coords) - 1; i >= 0; i-- {
				path.Coords = append(path.Coords, targetPath.Coords[i])
			}
		}

		if !contains(visited, current) {
			log.Printf("have not already visited: %v", current)
			visited = append(visited, current)
		}

		// Get the neighbours of the current node
		neighbours := GetNeighbourCoordsNOTWRAPPED(current.Coord, gs)

		searchNodeNeighbours := []SearchNode{}

		for _, neighbour := range neighbours {
			sn := SearchNode{Coord: neighbour, Parent: &current}
			if !CheckIfCoordIsSnake(sn.Coord, gs) {
				searchNodeNeighbours = append(searchNodeNeighbours, sn)
			}

		}

		for _, neighbour := range searchNodeNeighbours {
			if !contains(visited, neighbour) {
				queue = append(queue, neighbour)
			}
		}

	}

	return path
}

func contains(snl []SearchNode, sn SearchNode) bool {
	for _, v := range snl {
		if v.Coord == sn.Coord {
			return true
		}
	}
	return false
}
