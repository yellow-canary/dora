package utils

import (
	"context"

	"github.com/google/go-github/v52/github"
)

// getAllPullRequests retrieves all the pull requests of a GitHub repository
func GetAllPullRequests(client *github.Client, owner, repo string) ([]*github.PullRequest, error) {
	opt := &github.PullRequestListOptions{State: "all", ListOptions: github.ListOptions{PerPage: 100}}
	var pullRequests []*github.PullRequest

	for {
		prs, resp, err := client.PullRequests.List(context.Background(), owner, repo, opt)
		if err != nil {
			return nil, err
		}
		pullRequests = append(pullRequests, prs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return pullRequests, nil
}
