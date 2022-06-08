package graph

import (
	"search-alogs/constraints"
	"search-alogs/node"
	"search-alogs/queues"
	"sync"
)

// NewWithCost creates a new instance of the graph data structure,
// which has a 'cost'
func NewWithCost[T constraints.Key]() *WithCost[T] {
	return &WithCost[T]{
		Nodes:      node.NodesWithCost[T]{},
		Neighbours: make(map[node.NodeWithCost[T]][]*node.NodeWithCost[T]),
		Visited:    Set[T]{},
		Queue:      &queues.PriorityQueue[T]{},
		Steps:      make([][]T, 0),
		lock:       sync.RWMutex{},
	}
}

// Set is a unique 'Set' data structure, which takes a type argument
// for the key of the set
type Set[T constraints.Key] struct {
	Values map[T]bool
}

// Neighbours is a map of associated nodes
type Neighbours[T constraints.Key] map[node.NodeWithCost[T]][]*node.NodeWithCost[T]

// WithCost is a generic graph data structure, which takes a list of nodes
// and a list of each nodes neighbours. It's possible to use this graph
// data structure across multiple search algorithms
type WithCost[T constraints.Key] struct {
	Nodes      node.NodesWithCost[T]
	Neighbours map[node.NodeWithCost[T]][]*node.NodeWithCost[T]
	Visited    Set[T]
	Queue      *queues.PriorityQueue[T]
	lock       sync.RWMutex
	Steps      [][]T
}

// AddNode adds a new node to the graph
func (g *WithCost[T]) AddNode(newNode *node.NodeWithCost[T]) *WithCost[T] {
	if g.Nodes == nil {
		g.Nodes = make(node.NodesWithCost[T], 0)
	}

	g.Nodes = append(g.Nodes, newNode)

	return g
}

// AddNeighbour -
func (g *WithCost[T]) AddNeighbour(a, b *node.NodeWithCost[T]) *WithCost[T] {
	if g.Neighbours == nil {
		g.Neighbours = make(map[node.NodeWithCost[T]][]*node.NodeWithCost[T])
	}

	g.Neighbours[*a] = append(g.Neighbours[*a], b)
	g.Neighbours[*b] = append(g.Neighbours[*b], a)

	return g
}

// GetNeighbours -
func (g *WithCost[T]) GetNeighbours() map[T][]T {
	neighbours := make(map[T][]T)

	for k, v := range g.Neighbours {
		var nNeighbours []T
		for _, n := range v {
			nNeighbours = append(nNeighbours, n.Name)
		}

		neighbours[k.Name] = nNeighbours
	}

	return neighbours
}

// ListSteps lists the steps taken, or the list of frontiers
func (g *WithCost[T]) ListSteps() [][]T {
	return g.Steps
}
