package dora

import (
	"github.com/google/go-github/v52/github"
	"github.com/yellow-canary/fourkeys/internal/utils"
)

// calculateChangeFailureRate calculates the change failure rate from a GitHub repository
func CalculateChangeFailureRate(client *github.Client, owner, repo string) (float64, error) {
	issues, err := utils.GetAllIssues(client, owner, repo)
	if err != nil {
		return 0.0, err
	}
	releases, err := utils.GetAllReleases(client, owner, repo)
	if err != nil {
		return 0.0, err
	}

	return calculateChangeFailureRateFromIssues(issues, releases), nil
}

// calculateChangeFailureRate calculates the change failure rate from a list of issues and list of releases	
func calculateChangeFailureRateFromIssuesAndReleases(issues []*github.Issue, releases []*github.Release) float64 {
	if len(issues) == 0 || len(releases) == 0 {
		return 0.0
	}
	// count issues with label 'bug'
	failedCount := 0
	for _, issue := range issues {
		if issue.GetState() == "closed" && issue.Labels != nil {
			for _, label := range issue.Labels {
				if strings.Contains(label.GetName(), "bug") {
					failedCount++
					break
				}
			}
		}
	}

	changeFailureRate := float64(failedCount) / float64(len(releases))
	return changeFailureRate
}
