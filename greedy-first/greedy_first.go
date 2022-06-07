package greedy_first

import (
	"fmt"
	"log"
	"search-alogs/constraints"
	"search-alogs/graph"
	"search-alogs/node"
	"search-alogs/queues"
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

type GreedySearch[T constraints.Key] struct {
	*graph.WithCost[T]
}

func (g *GreedySearch[T]) String() {
	output := ""
	for i := 0; i < len(g.Nodes); i++ {
		output += g.Nodes[i].String() + " -> "
		near := g.Neighbours[*g.Nodes[i]]

		for j := 0; j < len(near); j++ {
			output += near[j].String() + " "
		}

		output += "\n"
	}

	// We ignore the linter here, because we want to
	// print the raw string data
	fmt.Println(output) //nolint:forbidigo
}

// Do takes a source (starting point), a target (endpoint)
func (g *GreedySearch[T]) Do(source, target *node.NodeWithCost[T]) []T {
	path := make([]T, 0)

	// Create a Set, this will house the 'visited' nodes, or 'closed' nodes.
	// Closed nodes are the nodes that have been visited and selected as the chosen path... I think?
	g.Visited = graph.Set[T]{Values: make(map[T]bool)}

	// Creates a priority queue, the order is decided by the 'cost' value. For instance,
	// the distance between two nodes. The shortest node is selected
	g.Queue = queues.NewPriorityQueue(g.Neighbours[*source])

	// log.Println("queue:", g.queue)

	// We add the source node as the first node in the queue
	g.Queue.Add(source)

	// We then also set the source node as 'visited' in the set
	g.Visited.Values[source.Name] = true

	// Next, we loop over the queue, whilst there are values
	// still in the queue
	for {
		if len(g.Queue.Nodes) > 0 {
			popped := g.Queue.Pop()

			frontier := g.Queue.List()
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
			for _, v := range g.Neighbours[*popped] {
				if g.Visited.Values[v.Name] == false {
					g.Visited.Values[v.Name] = true
					g.Queue.Add(v)
				}
			}
		}
	}
}
