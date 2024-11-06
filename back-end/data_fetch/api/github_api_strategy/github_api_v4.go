package github_api_strategy

import (
	"context"
	"data_fetch/pkg/utils"
	"data_fetch/repository/model"

	"github.com/shurcooL/githubv4"
)

type GitHubAPIV4Strategy struct{}

var clientV4List []*githubv4.Client

func (g *GitHubAPIV4Strategy) Init() {
	tcList := generateTCList()
	if tcList == nil {
		utils.GetLogger().Fatal("GitHub API V4 client initialized failed")
	}
	for _, tc := range tcList {
		clientV4List = append(clientV4List, githubv4.NewClient(tc))
	}
	utils.GetLogger().Info("GitHub API V4 client initialized")
}

func (g *GitHubAPIV4Strategy) Fetch() {
	variables := map[string]interface{}{
		"query": githubv4.String("user:*"), // 使用通配符搜索所有用户
		"first": githubv4.Int(100),         // 每页获取100个用户
		"after": githubv4.String(""),       // 初始游标为空
	}

	q := model.UserQuery{}
	err := clientV4List[0].Query(context.Background(), &q, variables)
	if err != nil {
		utils.GetLogger().Fatalf("Query failed: %v", err)
	}

	for _, user := range q.Search.Nodes {

		//todo Process the user data
		utils.GetLogger().Infof("process User: %v", user)
		// ... and other fields
	}

	//for {
	//	q := &UserQuery{}
	//	err := clientV4.Query(context.Background(), q, variables)
	//	if err != nil {
	//		utils.GetLogger().Fatalf("Query failed: %v", err)
	//	}
	//
	//	for _, user := range q.Search.Nodes {
	//
	//		//todo Process the user data
	//		utils.GetLogger().Infof("process User: %v", user)
	//		// ... and other fields
	//	}
	//
	//	if !q.Search.PageInfo.HasNextPage {
	//		break
	//	}
	//
	//	variables["cursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
	//}
}
