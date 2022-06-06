package greedy_first

import (
	"fmt"
	"log"
	"sync"
)

/* Target graph
                                                        3
										       A(5) -- -- -- C(8)
										      / 3             |
									         B(6)             | 3
									         | 2              |
							                 |     4       1  |
						                     D(4) -- E(4) --- F(5)
		                                               |     /
		                                             2 |    / 3
	                                                   |   /
                                                       |  /
                                         H(0) -- -- -- G(2)
*/

// Key is a generic type constraint, as we want to accept numbers
// or strings as a Key to our data structures
type Key interface {
	~string | ~int
}

// Set is a unique 'Set' data structure, which takes a type argument
// for the key of the set
type Set[T Key] struct {
	values map[T]bool
}

// Graph is a generic graph data structure, which takes a list of nodes
// and a list of each nodes neighbours. It's possible to use this graph
// data structure across multiple search algorithms
type Graph[T Key] struct {
	nodes      Nodes[T]
	neighbours map[Node[T]][]*Node[T]
	visited    Set[T]
	queue      *PriorityQueue[T]
	lock       sync.RWMutex
}

// NewGraph creates a new instance of the graph data structure
func NewGraph[T Key]() *Graph[T] {
	return &Graph[T]{
		nodes:      Nodes[T]{},
		neighbours: make(map[Node[T]][]*Node[T]),
		visited:    Set[T]{},
		queue:      &PriorityQueue[T]{},
		lock:       sync.RWMutex{},
	}
}

// AddNode adds a new node to the graph
func (g *Graph[T]) AddNode(node *Node[T]) *Graph[T] {
	if g.nodes == nil {
		g.nodes = make(Nodes[T], 0)
	}

	g.nodes = append(g.nodes, node)

	return g
}

// AddNeighbour -
func (g *Graph[T]) AddNeighbour(a, b *Node[T]) *Graph[T] {
	if g.neighbours == nil {
		g.neighbours = make(map[Node[T]][]*Node[T])
	}

	g.neighbours[*a] = append(g.neighbours[*a], b)
	g.neighbours[*b] = append(g.neighbours[*b], a)

	return g
}

func (g *Graph[T]) String() {
	output := ""
	for i := 0; i < len(g.nodes); i++ {
		output += g.nodes[i].String() + " -> "
		near := g.neighbours[*g.nodes[i]]

		for j := 0; j < len(near); j++ {
			output += near[j].String() + " "
		}

		output += "\n"
	}

	// We ignore the linter here, because we want to
	// print the raw string data
	fmt.Println(output) //nolint:forbidigo
}

// GreedySearch takes a source (starting point), a target (endpoint)
func (g *Graph[T]) GreedySearch(source, target *Node[T]) []T {
	path := make([]T, 0)

	// Create a Set, this will house the 'visited' nodes, or 'closed' nodes.
	// Closed nodes are the nodes that have been visited and selected as the chosen path... I think?
	g.visited = Set[T]{values: make(map[T]bool)}

	// Creates a priority queue, the order is decided by the 'cost' value. For instance,
	// the distance between two nodes. The shortest node is selected
	g.queue = NewQueue(g.neighbours[*source])

	// log.Println("queue:", g.queue)

	// We add the source node as the first node in the queue
	g.queue.Add(source)

	// We then also set the source node as 'visited' in the set
	g.visited.values[source.Name] = true

	// Next, we loop over the queue, whilst there are values
	// still in the queue
	for {
		if len(g.queue.Nodes) > 0 {
			popped := g.queue.Pop()

			frontier := g.queue.List()
			log.Println("frontier:", frontier)

			path = append(path, popped.Name)

			// If we reach the target, then WE ARE DONE 🎉
			if popped == target {
				return path
			}

			// We expand the neighbours of the next selected node
			// from the queue. We check that that node hasn't already
			// been visited. We then add it to the visited nodes,
			// and add the neighbours of that node to the queue.
			// This is similar to breadth first search, except that
			// the 'cost' value is considered as part of the queue.
			for _, v := range g.neighbours[*popped] {
				if g.visited.values[v.Name] == false {
					g.visited.values[v.Name] = true
					g.queue.Add(v)
				}
			}
		}
	}
}
