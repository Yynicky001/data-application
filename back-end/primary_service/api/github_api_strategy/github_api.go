package github_api_strategy

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/neo4j"
	"github.com/google/go-github/v66/github"
	"sync"
)

var (
	clientList    []*github.Client
	wg            *sync.WaitGroup                      // 用于等待所有goroutine完成
	maxGoroutines = 20                                 // 根据GitHub API限制调整
	semaphore     = make(chan struct{}, maxGoroutines) // 使用带缓冲的channel来控制并发
	usersBatch    []map[string]interface{}             // 用于存储批量数据
	batchSize     = 1000                               //批量存储到数据库的大小
)

type GitHubAPIDefaultStrategy struct{}

// Fetch GitHubAPI拉取策略
func (g *GitHubAPIDefaultStrategy) Fetch() {
	fetchUserData()
}

func fetchUserDetails(user *github.User) {
	wg = &sync.WaitGroup{}

}

//func fetchRepoData() {
//	wg = &sync.WaitGroup{}
//	index := 0
//	client := clientList[index]
//	// 获取仓库列表
//	opt := &github.RepositoryListAllOptions{
//		Since: 0,
//	}
//	for {
//		all, response, err := client.Repositories.ListAll(context.Background(), opt)
//		if err != nil {
//			utils.GetLogger().Errorf("Error fetching repositories: %v", err)
//		}
//
//	}
//}

// 拉取用户简略数据
func fetchUserData() {
	wg = &sync.WaitGroup{}
	var perPage = 500
	opts := &github.UserListOptions{ListOptions: github.ListOptions{PerPage: perPage}}
	opts.Page = 1
	opts.Since = 0
	index := 0
	client := clientList[index]
	for {
		users, response, err := client.Users.ListAll(context.Background(), opts)

		if err != nil {
			utils.GetLogger().Errorf("Error fetching users: %v", err)
		}
		utils.GetLogger().Infof("Fetching users since ID: %d", opts.Since)
		for _, user := range users {
			wg.Add(1)
			go fetchUsersConcurrently(user)
		}

		if response.NextPage == 0 || opts.Since >= 1000000 {
			utils.GetLogger().Infof("%+v", response)
			utils.GetLogger().Infof("All users fetched.")
			break
		}

		if response.StatusCode == 304 {
			index++
			if index >= len(clientList) {
				utils.GetLogger().Warnf("No more clients available. Exiting...")
				break
			}
			client = clientList[index]
			utils.GetLogger().Infof("Switched to client %d", index)
		}
		opts.Since = users[len(users)-1].GetID()
	}
	wg.Wait() // 等待所有goroutines完成
	if len(usersBatch) > 0 {
		if err := neo4j.BatchCreateUserNodes(usersBatch); err != nil {
			utils.GetLogger().Errorf("Error batch inserting users: %v", err)
		}
	}
}

// 并发拉取数据
func fetchUsersConcurrently(user *github.User) {
	defer wg.Done()
	semaphore <- struct{}{}        // 获取信号量
	defer func() { <-semaphore }() // 释放信号量

	utils.GetLogger().Infof("user data is processing: %s\n", user.GetLogin())
	usersBatch = append(usersBatch, neo4j.User2Map(user))
	if len(usersBatch) >= batchSize {
		if err := neo4j.BatchCreateUserNodes(usersBatch); err != nil {
			utils.GetLogger().Errorf("Error batch inserting users: %v", err)
		}
		utils.GetLogger().Infof("Batch insert completed. Size: %d", batchSize)
		usersBatch = make([]map[string]interface{}, 0) // 清空切片
	}
}

// Init 初始化GitHub客户端
func (g *GitHubAPIDefaultStrategy) Init() {
	// 生成tc list
	tcList := generateTCList()
	// 创建OAuth2客户端
	for _, tc := range tcList {
		clientList = append(clientList, github.NewClient(tc))
	}

	utils.GetLogger().Info("GitHub API client initialized")
}
