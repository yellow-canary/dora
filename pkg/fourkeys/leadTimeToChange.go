package fourkeys

import (
	"time"

	"github.com/google/go-github/v52/github"
	"github.com/yellow-canary/fourkeys/internal/utils"
)

// calculateLeadTimeToChange calculates the lead time to change from a GitHub repository
func CalculateLeadTimeToChange(client *github.Client, owner, repo string) (time.Duration, error) {
	pullRequests, err := utils.GetAllPullRequests(client, owner, repo)
	if err != nil {
		return 0, err
	}

	return calculateLeadTimeToChangeFromPullRequests(pullRequests), nil
}

// calculateLeadTimeToChange calculates the lead time to change from a list of pull requests
func calculateLeadTimeToChangeFromPullRequests(pullRequests []*github.PullRequest) time.Duration {
	if len(pullRequests) == 0 {
		return 0
	}

	var totalTime time.Duration
	for _, pr := range pullRequests {
		if pr.MergedAt == nil || pr.CreatedAt == nil {
			continue
		}
		duration := pr.MergedAt.Sub(pr.CreatedAt.Time)
		totalTime += duration
	}

	// Calculate the average lead time to change
	averageLeadTime := totalTime / time.Duration(len(pullRequests))
	return averageLeadTime
}
