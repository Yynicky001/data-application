package main

import (
	strategy "github-data-evaluator/api/github_api_strategy"
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/router"
)

func main() {
	utils.GetLogger().Info("akkk")
	//serverStart()
	//fetchStart()
}

func serverStart() {
	r := router.NewRouter()
	_ = r.Run(config.Conf.Server.Port)
}

func fetchData() {
	context := &strategy.GitHubAPIContext{}

	if config.Conf.GitHub.Strategy == "v4" {
		context.SetGitHubAPIContext(&strategy.GitHubAPIV4Strategy{})
	} else if config.Conf.GitHub.Strategy == "default" {
		context.SetGitHubAPIContext(&strategy.GitHubAPIDefaultStrategy{})
	} else {
		panic("unknown github api strategy")
	}

	context.Fetch()
}
