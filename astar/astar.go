package astar

/*
 'A': 'C', 'D'
*/

type Cost map[string]int

type Node struct {
	Cost       Cost
	Neighbours Graph
}

type Graph map[string]Node

type AStar struct {
	Graph Graph
}

// New a start search algorithm
func New(graph Graph) *AStar {
	return &AStar{graph}
}
