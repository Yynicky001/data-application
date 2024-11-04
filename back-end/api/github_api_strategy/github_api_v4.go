package github_api_strategy

import (
	"context"
	"fmt"
	"github-data-evaluator/config"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubAPIV4Strategy struct{}

var clientV4 *githubv4.Client

func (g *GitHubAPIV4Strategy) Init() {
	GithubToken := config.Conf.GitHub.Token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	clientV4 = githubv4.NewClient(tc)
}

func (g *GitHubAPIV4Strategy) Fetch() {
	q := &Query{}
	err := clientV4.Query(context.Background(), q, nil)
	if err != nil {
		panic(err)
	}

	println(q.Viewer.ID)
	fmt.Printf("%+v\n\n", q.Viewer)
}

type Query struct {
	Viewer Viewer `graphql:"viewer"`
}

type Viewer struct {
	Login     string
	ID        int64
	AvatarURL string
	HTMLURL   string            `graphql:"url"`
	CreatedAt githubv4.DateTime `graphql:"createdAt"`
	Company   string
	Location  string
	Followers struct {
		TotalCount int
	} `graphql:"followers"`
	Followings struct {
		TotalCount int
	} `graphql:"following"`
	Repositories struct {
		Edges []struct {
			Node RepoGQL
		}
	}
	FollowingInfos struct {
		Edges []struct {
			Node UserGQL
		}
	}
	FollowerInfos struct {
		Edges []struct {
			Node UserGQL
		}
	}
}

type RepoGQL struct {
	Name      string
	Owner     string `graphql:"nameWithOwner"`
	Languages struct {
		Color string
		Name  string
	} `graphql:"primaryLanguage"`
	starCount     int                      `graphql:"stargazerCount"`
	WatcherCount  struct{ TotalCount int } `graphql:"watchers"`
	ForkCount     int                      `graphql:"forkCount"`
	HTMLURL       string                   `graphql:"url"`
	CreatedAt     githubv4.DateTime
	UpdatedAt     githubv4.DateTime
	Collaborators struct {
		TotalCount int
		Edges      []struct {
			Node struct {
				CommitCount      int
				PullRequestCount struct{ TotalCount int }
			}
		}
	}
}

type UserGQL struct {
	Login     string
	ID        string
	AvatarURL string
	Followers struct{ TotalCount int }
	Following struct{ TotalCount int }
	CreatedAt githubv4.DateTime
	URL       string
	Company   string
	Location  string
}
