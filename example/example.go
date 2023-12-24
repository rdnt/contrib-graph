package main

import (
	"os"

	graph "github.com/rdnt/contributions-graph"
)

func main() {
	contribs := []graph.ContributionDay{
		{Count: 0, Color: "#ebedf0"},
		{Count: 1, Color: "#ebedf0"},
		{Count: 2, Color: "#ebedf0"},
		{Count: 3, Color: "#ebedf0"},
		{Count: 4, Color: "#ebedf0"},
	}

	g := graph.NewGraph(contribs)

	fd, err := os.Create("contributions-dark-halloween.svg")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	err = g.Render(fd, graph.Dark, true)
	if err != nil {
		panic(err)
	}

	fl, err := os.Create("contributions-light.svg")
	if err != nil {
		panic(err)
	}
	defer fl.Close()

	err = g.Render(fl, graph.Light, false)
	if err != nil {
		panic(err)
	}
}
