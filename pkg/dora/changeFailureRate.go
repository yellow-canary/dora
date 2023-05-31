package dora

import (
	"github.com/google/go-github/v52/github"
	"github.com/yellow-canary/dora/internal/utils"
)

// calculateChangeFailureRate calculates the change failure rate from a GitHub repository
func CalculateChangeFailureRate(client *github.Client, owner, repo string) (float64, error) {
	issues, err := utils.GetAllIssues(client, owner, repo)
	if err != nil {
		return 0.0, err
	}

	return calculateChangeFailureRateFromIssues(issues), nil
}

// calculateChangeFailureRate calculates the change failure rate from a list of issues
func calculateChangeFailureRateFromIssues(issues []*github.Issue) float64 {
	if len(issues) == 0 {
		return 0.0
	}

	failedCount := 0
	for _, issue := range issues {
		if issue.GetState() == "closed" && issue.Labels != nil {
			for _, label := range issue.Labels {
				if label.GetName() == "bug" {
					failedCount++
					break
				}
			}
		}
	}

	changeFailureRate := float64(failedCount) / float64(len(issues))
	return changeFailureRate
}
