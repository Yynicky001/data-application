package github_api_strategy

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
