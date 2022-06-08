package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"search-alogs/graph"
	greedy_first "search-alogs/greedy-first"
	"search-alogs/node"
)

func main() {
	engine := html.New("./views", ".html.tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main.html.tmpl",
	})

	app.Get("/greedy-best-first", func(ctx *fiber.Ctx) error {
		return ctx.Render("greedy-best-first", nil)
	})

	app.Get("/data/greedy-best-first", func(ctx *fiber.Ctx) error {
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

		searcher := greedy_first.GreedySearch[string]{g}
		path := searcher.Do(a5, h0)
		frontiers := searcher.ListSteps()

		return ctx.JSON(map[string]interface{}{
			"path":       path,
			"frontiers":  frontiers,
			"nodes":      g.Nodes,
			"neighbours": g.GetNeighbours(),
		})
	})

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}
}
