package utils

import (
	"context"

	"github.com/google/go-github/v52/github"
)

// getAllIssues retrieves all the issues of a GitHub repository
func GetAllIssues(client *github.Client, owner, repo string) ([]*github.Issue, error) {
	opt := &github.IssueListByRepoOptions{State: "all", ListOptions: github.ListOptions{PerPage: 100}}
	var issues []*github.Issue

	for {
		issueList, resp, err := client.Issues.ListByRepo(context.Background(), owner, repo, opt)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issueList...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return issues, nil
}
