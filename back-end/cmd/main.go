package main

import (
	strategy "github-data-evaluator/api/github_api_strategy"
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/router"
)

func main() {

	if f, exists := commandMap[config.Conf.Server.Command]; exists {
		f()
	} else {
		utils.GetLogger().Fatal("unknown environment")
	}
}

var commandMap = map[string]func(){
	"server": serverStart,
	"fetch":  fetchData,
}

func serverStart() {
	r := router.NewRouter()
	_ = r.Run(config.Conf.Server.Port)
}

func fetchData() {
	c := &strategy.GitHubAPIContext{}
	switch config.Conf.GitHub.Strategy {
	case "v4":
		c.SetGitHubAPIContext(&strategy.GitHubAPIV4Strategy{})
	default:
		c.SetGitHubAPIContext(&strategy.GitHubAPIDefaultStrategy{})
	}
	c.Fetch()
}
