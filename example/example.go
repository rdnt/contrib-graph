package main

import (
	"os"

	graph "github.com/rdnt/contributions-graph"
)

func main() {
	contribs := []graph.ContributionDay{
		{Count: 0, Color: graph.Color0},
		{Count: 1, Color: graph.Color1},
		{Count: 2, Color: graph.Color2},
		{Count: 3, Color: graph.Color3},
		{Count: 4, Color: graph.Color4},
	}

	g := graph.NewGraph(contribs)

	fd, err := os.Create("contributions-dark-halloween.svg")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	err = g.Render(fd, graph.HalloweenDarkTheme)
	if err != nil {
		panic(err)
	}

	fl, err := os.Create("contributions-light.svg")
	if err != nil {
		panic(err)
	}
	defer fl.Close()

	err = g.Render(fl, graph.LightTheme)
	if err != nil {
		panic(err)
	}
}
