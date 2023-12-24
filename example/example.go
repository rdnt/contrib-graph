package main

import (
	"os"

	graph "github.com/rdnt/contributions-graph"
)

func main() {
	contribs := []graph.ContributionDay{
		//{Count: 0, Color: "#ebedf0"},
		//{Count: 1, Color: "#da56dc"},
		//{Count: 2, Color: "#dc56da"},
		//{Count: 3, Color: "#dc56d3"},
		//{Count: 4, Color: "#dc56b6"},
		{Count: 0, Color: graph.Color0},
		{Count: 1, Color: graph.Color1},
		{Count: 2, Color: graph.Color2},
		{Count: 3, Color: graph.Color3},
		{Count: 4, Color: graph.Color4},
	}

	g := graph.NewGraph(contribs)

	//fd, err := os.Create("contributions-dark-halloween.svg")
	//if err != nil {
	//	panic(err)
	//}
	//defer fd.Close()
	//
	//err = g.Render(fd, graph.Dark, graph.Halloween)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fl, err := os.Create("contributions-light.svg")
	//if err != nil {
	//	panic(err)
	//}
	//defer fl.Close()
	//
	//err = g.Render(fl, graph.Light, graph.Normal)
	//if err != nil {
	//	panic(err)
	//}

	fo, err := os.Create("old.svg")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	err = g.Render(fo, graph.DarkHalloween)
	if err != nil {
		panic(err)
	}

	fn, err := os.Create("new.svg")
	if err != nil {
		panic(err)
	}
	defer fn.Close()

	err = g.Render(fn, graph.Custom)
	if err != nil {
		panic(err)
	}
}
