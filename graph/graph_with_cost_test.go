package graph

import (
	"github.com/matryer/is"
	"search-alogs/node"
	"testing"
)

var (
	s0 = &node.NodeWithCost[string]{
		Name: "B",
		Cost: 6,
	}

	c3 = &node.NodeWithCost[string]{
		Name: "C",
		Cost: 3,
	}

	a10 = &node.NodeWithCost[string]{
		Name: "A",
		Cost: 10,
	}
)

func TestCanAddNode(t *testing.T) {
	is := is.New(t)
	g := NewWithCost[string]()
	g = g.AddNode(s0).AddNode(c3).AddNeighbour(s0, c3)
	is.Equal(len(g.Nodes), 2)
	is.Equal(len(g.Neighbours), 2)
}
