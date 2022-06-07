package queues

import (
	"search-alogs/constraints"
	"search-alogs/node"
	"sort"
)

// PriorityQueue -
type PriorityQueue[T constraints.Key] struct {
	Nodes node.NodesWithCost[T]
}

// NewPriorityQueue instance
func NewPriorityQueue[T constraints.Key](nodes node.NodesWithCost[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{nodes}
}

// Add node
func (q *PriorityQueue[T]) Add(node *node.NodeWithCost[T]) node.NodesWithCost[T] {
	q.Nodes = append(q.Nodes, node)
	sort.Sort(q.Nodes)

	return q.Nodes
}

// Pop -
func (q *PriorityQueue[T]) Pop() *node.NodeWithCost[T] {
	lenNodes := len(q.Nodes)
	tail := q.Nodes[1:lenNodes]
	head := q.Nodes[0]

	q.Nodes = tail

	return head
}

// List values in the queue
func (q *PriorityQueue[T]) List() node.NodesWithCost[T] {
	return q.Nodes
}
