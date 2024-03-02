package github

import (
	"context"
	"time"

	"github.com/samber/lo"

	"github.com/rdnt/contribs-graph/graph"
)

type Client interface {
	GetContributions(ctx context.Context, user string, from, to time.Time) (ContributionsResponse, error)
}

type ContributionsResponse struct {
	IsHalloween   bool
	Contributions []Contribution
}

type Contribution struct {
	Count int
	Color string
}

func Contributions(ctx context.Context, client Client, user string) ([]graph.ContributionDay, error) {
	weeksInYear := 52
	fromOffset := (weeksInYear-1)*7 + int(time.Now().Weekday())
	from := time.Now().AddDate(0, 0, -fromOffset)
	to := time.Now()

	contribsView, err := client.GetContributions(ctx, user, from, to)
	if err != nil {
		return nil, err
	}

	contribs := lo.Map(contribsView.Contributions, func(c Contribution, _ int) graph.ContributionDay {
		return graph.ContributionDay(c)
	})

	return contribs, nil
}
