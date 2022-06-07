package queues

import (
	"github.com/matryer/is"
	"search-alogs/node"
	"testing"
)

func TestAddToQueue(t *testing.T) {
	is := is.New(t)
	q := NewPriorityQueue[string]()
	_ = q.Add(&node.NodeWithCost[string]{
		Name: "a",
		Cost: 1,
	})

	_ = q.Add(&node.NodeWithCost[string]{
		Name: "b",
		Cost: 2,
	})

	values := q.List()

	is.Equal(len(values), 2)

	node := q.Pop()
	is.Equal(node.Cost, 1)
	is.Equal(node.Name, "a")

	is.Equal(len(q.List()), 1)

	node = q.Pop()
	is.Equal(node.Cost, 2)
	is.Equal(node.Name, "b")

	is.Equal(len(q.List()), 0)
}
