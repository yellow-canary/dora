package dora

import (
	"math"

	"github.com/google/go-github/v52/github"
	"github.com/yellow-canary/fourkeys/internal/utils"
)

// calculateDeploymentFrequency calculates the deployment frequency from a GitHub repository
func CalculateDeploymentFrequency(client *github.Client, owner, repo string) (float64, error) {
	releases, err := utils.GetAllReleases(client, owner, repo)
	if err != nil {
		return 0.0, err
	}

	return calculateDeploymentFrequencyFromReleases(releases), nil
}

// calculateDeploymentFrequency calculates the deployment frequency from a list of releases
func calculateDeploymentFrequencyFromReleases(releases []*github.RepositoryRelease) float64 {
	if len(releases) <= 1 {
		return 0.0
	}

	var totalDays float64
	for i := 1; i < len(releases); i++ {
		prevRelease := releases[i-1]
		currentRelease := releases[i]

		// Calculate the time between two releases in days
		duration := currentRelease.GetPublishedAt().Sub(prevRelease.GetPublishedAt().Time)
		totalDays += duration.Hours() / 24
	}

	// Calculate the deployment frequency in releases per week
	deploymentFrequency := math.Abs(float64(len(releases)-1) / (totalDays / 7))
	return deploymentFrequency
}
