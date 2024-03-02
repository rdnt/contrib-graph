package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/rdnt/contributions-graph/example/githubql"
	"github.com/rdnt/contributions-graph/github"
	"github.com/rdnt/contributions-graph/graph"
)

func main() {
	accessToken := os.Getenv("ACCESS_TOKEN")
	username := os.Getenv("USERNAME")

	tokSrc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})

	githubClient := githubql.New(
		"https://api.github.com/graphql",
		oauth2.NewClient(context.Background(), tokSrc),
	)

	contribs, err := github.Contributions(context.Background(), githubClient, username)
	handleErr(err)

	f, err := os.Create("contributions.svg")
	handleErr(err)
	defer f.Close()

	g := graph.New(contribs)

	err = g.Render(f, graph.DarkTheme)
	handleErr(err)

	log.Println("Saved to contributions.svg")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
