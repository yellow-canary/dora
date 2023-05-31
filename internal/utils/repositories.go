package utils

import (
	"context"
	"log"

	"github.com/google/go-github/v52/github"
)

func GetRepositoriesFullNames(ctx context.Context, client *github.Client) []string {
	var repoFullNames []string
	// Fetch the repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	if err != nil {
		log.Fatal("Error fetching repositories:", err)
	}
	// Print the full names of the repositories
	for _, repo := range repos {
		repoFullNames = append(repoFullNames, repo.GetFullName())
	}
	return repoFullNames
}
