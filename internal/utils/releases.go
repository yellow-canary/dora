package utils

import (
	"context"

	"github.com/google/go-github/v52/github"
)

// getAllReleases retrieves all the releases of a GitHub repository
func GetAllReleases(client *github.Client, owner, repo string) ([]*github.RepositoryRelease, error) {
	opt := &github.ListOptions{PerPage: 100}
	var releases []*github.RepositoryRelease

	for {
		repoReleases, resp, err := client.Repositories.ListReleases(context.Background(), owner, repo, opt)
		if err != nil {
			return nil, err
		}
		releases = append(releases, repoReleases...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return releases, nil
}
