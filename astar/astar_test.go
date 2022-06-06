package astar

import (
	"github.com/matryer/is"
	"testing"
)

func TestNewGraph(t *testing.T) {
	is := is.New(t)
	graph := Graph{}

	/*
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

	graph["A"] = Node{
		Cost: 5,
		Neighbours: Graph{
			"B": Node{
				Neighbours: Graph{
					"D": Node{
						Neighbours: nil,
						Cost: Cost{
							"B": 1,
						},
					},
				},
			},
			"C": Node{
				Neighbours: nil,
			},
		},
	}

	// is.Equal(graph["A"].Neighbours["B"].Neighbours["D"].Costs["B"], "C")
}
