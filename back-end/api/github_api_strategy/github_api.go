package github_api_strategy

import (
	"context"
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/service"
	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
	"sync"
)

var (
	client          *github.Client
	wg              *sync.WaitGroup                      // 用于等待所有goroutine完成
	maxGoroutines   = 20                                 // 根据GitHub API限制调整
	semaphore       = make(chan struct{}, maxGoroutines) // 使用带缓冲的channel来控制并发
	developersBatch []*model.Developer
	batchSize       = 200 //批量存储到数据库的大小
)

type GitHubAPIDefaultStrategy struct{}

// Fetch GitHubAPI拉取策略
func (g *GitHubAPIDefaultStrategy) Fetch() {
	wg = &sync.WaitGroup{}
	var perPage = 100
	opts := &github.UserListOptions{ListOptions: github.ListOptions{PerPage: perPage}}
	opts.Page = 1
	opts.Since = 0
	for {
		users, response, err := client.Users.ListAll(context.Background(), opts)

		if err != nil {
			utils.GetLogger().Errorf("Error fetching users: %v", err)
		}
		utils.GetLogger().Infof("Fetching users since ID: %d", opts.Since)
		for _, user := range users {
			wg.Add(1)
			go fetchUserDetailsConcurrently(user)
		}
		if response.NextPage == 0 || response.StatusCode == 304 || opts.Since >= 100000 {
			utils.GetLogger().Infof("%+v", response)
			utils.GetLogger().Infof("All users fetched.")
			break
		}
		opts.Since = users[len(users)-1].GetID()
	}
	wg.Wait() // 等待所有goroutines完成
	if len(developersBatch) > 0 {
		service.GetDeveloperService().BatchInsertDevelopers(developersBatch)
	}
}

// 并发拉取数据
func fetchUserDetailsConcurrently(user *github.User) {
	defer wg.Done()
	semaphore <- struct{}{}        // 获取信号量
	defer func() { <-semaphore }() // 释放信号量

	utils.GetLogger().Infof("user data is processing: %s\n", user.GetLogin())
	developersBatch = append(developersBatch, model.User2Developer(user))
	if len(developersBatch) >= batchSize {
		service.GetDeveloperService().BatchInsertDevelopers(developersBatch)
		developersBatch = make([]*model.Developer, 0) // 清空切片
	}
}

// Init 初始化GitHub客户端
func (g *GitHubAPIDefaultStrategy) Init() {
	GithubToken := config.Conf.GitHub.Token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client = github.NewClient(tc)
	//client = github.NewClient(http.DefaultClient)
}
