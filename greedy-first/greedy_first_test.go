package greedy_first

import (
	"github.com/matryer/is"
	"testing"
)

var (
	s0 = &Node[string]{
		Name: "B",
		Cost: 6,
	}

	c3 = &Node[string]{
		Name: "C",
		Cost: 3,
	}

	a10 = &Node[string]{
		Name: "A",
		Cost: 10,
	}
)

func TestCanPioritiseQueue(t *testing.T) {
	is := is.New(t)

	var nodes []*Node[string]
	queue := NewQueue(nodes)

	updated := queue.Add(s0)
	updated = queue.Add(c3)

	is.Equal(updated[0].Cost, 3)
	is.Equal(updated[1].Cost, 6)
	is.Equal(updated[2].Cost, 10)
}

func TestCanAddNode(t *testing.T) {
	g := NewGraph[string]()
	g = g.AddNode(s0).AddNode(c3).AddNeighbour(s0, c3)
	g.String()
}

func TestGreedyFirstSearch(t *testing.T) {
	is := is.New(t)

	g := NewGraph[string]()

	a5 := NewNode("a", 5)
	b6 := NewNode("b", 6)
	c8 := NewNode("c", 8)
	d4 := NewNode("d", 4)
	e4 := NewNode("e", 4)
	f5 := NewNode("f", 5)
	g2 := NewNode("g", 2)
	h0 := NewNode("h", 0)

	g = g.
		AddNode(a5).AddNode(b6).AddNode(c8).AddNode(d4).
		AddNode(e4).AddNode(f5).AddNode(g2).AddNode(h0).
		AddNeighbour(a5, b6).AddNeighbour(a5, c8).AddNeighbour(b6, d4).AddNeighbour(c8, f5).
		AddNeighbour(d4, e4).AddNeighbour(e4, f5).AddNeighbour(e4, g2).AddNeighbour(f5, g2).
		AddNeighbour(g2, h0)

	path := g.GreedySearch(a5, h0)
	is.Equal(path, []string{"a", "b", "d", "e", "g", "h"})
}
