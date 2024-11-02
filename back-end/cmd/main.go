package main

import (
	"github-data-evaluator/api"
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/db"
)

func main() {
	loading()
	api.FetchDeveloperStart()
}

func loading() {
	utils.InitLog()
	config.InitConfig()
	db.InitDB()
	api.InitGithubClient()
}
