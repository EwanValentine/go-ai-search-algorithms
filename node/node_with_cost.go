package node

import (
	"fmt"
	"search-alogs/constraints"
)

// NewNodeWithCost -
func NewNodeWithCost[T constraints.Key](name T, cost int) *NodeWithCost[T] {
	return &NodeWithCost[T]{
		Name: name,
		Cost: cost,
	}
}

type NodesWithCost[T constraints.Key] []*NodeWithCost[T]

func (q NodesWithCost[T]) Len() int           { return len(q) }
func (q NodesWithCost[T]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q NodesWithCost[T]) Less(i, j int) bool { return q[i].Cost < q[j].Cost }

// NodeWithCost -
type NodeWithCost[T constraints.Key] struct {
	Name T
	Cost int
}

// String -
func (q *NodeWithCost[T]) String() string {
	return fmt.Sprintf("%v", q.Name)
}
