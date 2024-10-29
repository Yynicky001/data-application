package utils

import (
	"context"
	"github-data-evaluator/repository/model"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
	"sync"
)

var (
	client          *github.Client
	wg              *sync.WaitGroup // 用于等待所有goroutine完成
	maxGoroutines   = 20            // 根据GitHub API限制调整
	mutex           sync.Mutex
	visited         map[string]bool                      // 用于记录已访问的用户ID，防止重复查询
	semaphore       = make(chan struct{}, maxGoroutines) // 使用带缓冲的channel来控制并发
	developersBatch []*model.Developer
	batchSize       = 100 //批量存储到数据库的大小
)

func FetchStart() {
	wg = &sync.WaitGroup{}
	visited = make(map[string]bool)
	opts := &github.UserListOptions{Since: 0, ListOptions: github.ListOptions{PerPage: 10}}
	for {
		users, response, err := client.Users.ListAll(context.Background(), opts)

		if err != nil {
			LogrusObj.Error("Error fetching users: %v", err)
		}
		for _, user := range users {
			wg.Add(1)
			dfs(*user.Login)
		}
		if response.NextPage == 0 {
			wg.Wait() // 等待所有goroutines完成
			return
		}
	}

}

// 并发版本拉取数据
func fetchUserDetailsConcurrently(username string) {
	defer wg.Done()
	semaphore <- struct{}{}        // 获取信号量
	defer func() { <-semaphore }() // 释放信号量

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		LogrusObj.Error("Error fetching details for user %s: %v\n", username, err)
		return
	}

	LogrusObj.Infoln("Fetched details for user: %s", user.GetLogin())
	developer := &model.Developer{}
	developersBatch = append(developersBatch, developer.User2Developer(user))
	if len(developersBatch) >= batchSize {
		batchInsertDevelopers(developersBatch)
		developersBatch = nil
	}
}

// 批量存储开发者数据
func batchInsertDevelopers(batch []*model.Developer) {
	LogrusObj.Infoln("Inserted %d developers into the database.", len(batch))
	//todo 执行批量插入操作
}

// 深度优先搜索 查询用户数据
func dfs(username string) {
	if isVisited(username) {
		return
	}
	setVisited(username)

	go fetchUserDetailsConcurrently(username)

	opts := &github.ListOptions{PerPage: 10}

	for {
		// 获取当前页的followers列表
		followers, resp, err := client.Users.ListFollowers(context.Background(), username, opts)
		if err != nil {
			LogrusObj.Errorln("Error fetching following: %v", err)
		}

		// 遍历当前页的following
		for _, follower := range followers {
			wg.Add(1)
			go dfs(*follower.Login)
		}

		// 检查是否还有下一页
		if resp.NextPage == 0 {
			break
		}

		// 更新ListOptions的Page字段以获取下一页
		opts.Page = resp.NextPage
	}

	for {
		// 获取当前页的following列表
		followings, resp, err := client.Users.ListFollowing(context.Background(), username, opts)
		if err != nil {
			LogrusObj.Errorln("Error fetching following: %v", err)
		}

		// 遍历当前页的following
		for _, following := range followings {
			wg.Add(1)
			go fetchUserDetailsConcurrently(*following.Login)
		}

		// 检查是否还有下一页
		if resp.NextPage == 0 {
			break
		}

		// 更新ListOptions的Page字段以获取下一页
		opts.Page = resp.NextPage
	}
}

// InitGithubClient 初始化GitHub客户端
func InitGithubClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client = github.NewClient(tc)

	//client = github.NewClient(http.DefaultClient)
}

func setVisited(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	visited[key] = true
}

func isVisited(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	return visited[key]
}
