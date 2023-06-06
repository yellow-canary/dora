package fourkeys

import (
	"strings"
	"time"

	"github.com/google/go-github/v52/github"
	"github.com/yellow-canary/fourkeys/internal/utils"
)

// calculateTimeToRestoreService calculates the time to restore service from a GitHub repository
func CalculateTimeToRestoreService(client *github.Client, owner, repo string) (time.Duration, error) {
	issues, err := utils.GetAllIssues(client, owner, repo)
	if err != nil {
		return 0, err
	}

	return calculateTimeToRestoreServiceFromIssues(issues), nil
}

// calculateTimeToRestoreService calculates the time to restore service from a list of issues
func calculateTimeToRestoreServiceFromIssues(issues []*github.Issue) time.Duration {
	if len(issues) == 0 {
		return 0
	}

	var totalTime time.Duration
	var closedCount int
	for _, issue := range issues {
		if issue.GetState() == "closed" && issue.Labels != nil {
			for _, label := range issue.Labels {
				if strings.Contains(label.GetName(), "bug") {
					duration := issue.GetClosedAt().Sub(issue.GetCreatedAt().Time)
					totalTime += duration
					closedCount++
				}
			}
		}
	}
	if closedCount == 0 {
		return 0
	}

	// Calculate the average time to restore service
	averageTimeToRestore := totalTime / time.Duration(closedCount)
	return averageTimeToRestore
}
