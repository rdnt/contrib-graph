package githubql

import (
	"context"
	"net/http"
	"time"

	_ "github.com/Khan/genqlient/generate"
	"github.com/Khan/genqlient/graphql"

	graph "github.com/rdnt/contributions-graph/github"
)

//go:generate go run github.com/Khan/genqlient genql.yaml

type Client struct {
	apiURL   string
	gql      graphql.Client
	clientId string
}

func New(apiURL string, httpClient *http.Client) *Client {
	return &Client{
		apiURL: apiURL,
		gql:    graphql.NewClient(apiURL, httpClient),
	}
}

func (c *Client) GetContributions(ctx context.Context, user string, from, to time.Time) (graph.ContributionsResponse, error) {
	resp, err := contributionsView(ctx, c.gql, user, from, to)
	if err != nil {
		return graph.ContributionsResponse{}, err
	}

	isHaloween := resp.User.ContributionsCollection.ContributionCalendar.IsHalloween

	contribs := graph.ContributionsResponse{
		IsHalloween: isHaloween,
	}

	for _, w := range resp.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, d := range w.ContributionDays {
			contribs.Contributions = append(contribs.Contributions, graph.Contribution{
				Count: d.ContributionCount,
				Color: normalizeColor(d.Color, isHaloween),
			})
		}
	}

	return contribs, nil
}

func normalizeColor(color string, haloween bool) string {
	if !haloween {
		return color
	}

	switch color {
	case "#ebedf0":
		return "#ebedf0"
	case "#ffee4a":
		return "#9be9a8"
	case "#ffc501":
		return "#40c463"
	case "#fe9600":
		return "#30a14e"
	case "#03001c":
		return "#216e39"
	default:
		return "#ebedf0"
	}
}
