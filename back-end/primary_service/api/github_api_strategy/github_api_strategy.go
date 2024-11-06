package github_api_strategy

import (
	"context"
	"github-data-evaluator/config"
	"golang.org/x/oauth2"
	"net/http"
)

type GitHubFetchStrategy interface {
	Init()
	Fetch()
}

type GitHubAPIContext struct {
	strategy GitHubFetchStrategy
}

func (g *GitHubAPIContext) SetGitHubAPIContext(strategy GitHubFetchStrategy) {
	g.strategy = strategy
	g.strategy.Init()
}

func (g *GitHubAPIContext) Fetch() {
	g.strategy.Fetch()
}

func generateTCList() []*http.Client {
	GithubToken := config.Conf.GitHub.Token
	if len(GithubToken) == 0 {
		return nil
	}
	tcList := make([]*http.Client, 0)
	for _, token := range GithubToken {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tcList = append(tcList, oauth2.NewClient(context.Background(), ts))
	}
	return tcList
}
