package greedy_first

import (
	"fmt"
	"sort"
)

func NewNode[T Key](name T, cost int) *Node[T] {
	return &Node[T]{
		Name: name,
		Cost: cost,
	}
}

// Node -
type Node[T Key] struct {
	Name T
	Cost int
}

func (q *Node[T]) String() string {
	return fmt.Sprintf("%v", q.Name)
}

type PriorityQueue[T Key] struct {
	Nodes Nodes[T]
}

func NewQueue[T Key](nodes Nodes[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{nodes}
}

type Nodes[T Key] []*Node[T]

func (q Nodes[T]) Len() int           { return len(q) }
func (q Nodes[T]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q Nodes[T]) Less(i, j int) bool { return q[i].Cost < q[j].Cost }

// Add node
func (q *PriorityQueue[T]) Add(node *Node[T]) Nodes[T] {
	q.Nodes = append(q.Nodes, node)
	sort.Sort(q.Nodes)

	return q.Nodes
}

// Pop -
func (q *PriorityQueue[T]) Pop() *Node[T] {
	lenNodes := len(q.Nodes)
	tail := q.Nodes[1:lenNodes]
	head := q.Nodes[0]

	q.Nodes = tail

	return head
}

// List values in the queue
func (q *PriorityQueue[T]) List() Nodes[T] {
	return q.Nodes
}
