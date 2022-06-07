package greedy_first

import (
	"github.com/matryer/is"
	"search-alogs/graph"
	"search-alogs/node"
	"testing"
)

func TestGreedyFirstSearch(t *testing.T) {
	is := is.New(t)

	g := graph.NewWithCost[string]()

	a5 := node.NewNodeWithCost("a", 5)
	b6 := node.NewNodeWithCost("b", 6)
	c8 := node.NewNodeWithCost("c", 8)
	d4 := node.NewNodeWithCost("d", 4)
	e4 := node.NewNodeWithCost("e", 4)
	f5 := node.NewNodeWithCost("f", 5)
	g2 := node.NewNodeWithCost("g", 2)
	h0 := node.NewNodeWithCost("h", 0)

	g = g.
		AddNode(a5).AddNode(b6).AddNode(c8).AddNode(d4).
		AddNode(e4).AddNode(f5).AddNode(g2).AddNode(h0).
		AddNeighbour(a5, b6).AddNeighbour(a5, c8).AddNeighbour(b6, d4).AddNeighbour(c8, f5).
		AddNeighbour(d4, e4).AddNeighbour(e4, f5).AddNeighbour(e4, g2).AddNeighbour(f5, g2).
		AddNeighbour(g2, h0)

	searcher := GreedySearch[string]{g}
	path := searcher.Do(a5, h0)

	is.Equal(path, []string{"a", "b", "d", "e", "g", "h"})
}
