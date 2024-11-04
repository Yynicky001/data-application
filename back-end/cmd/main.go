package main

import (
	"github-data-evaluator/api"
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/db"
	"github-data-evaluator/router"
)

func main() {
	loading()
	serverStart()
	//api.FetchStart()
}

func serverStart() {
	r := router.NewRouter()
	_ = r.Run(config.Conf.Server.Port)
}

func loading() {
	utils.InitLog()
	config.InitConfig()
	db.InitDB()
	api.InitGithubClient()
	//api.InitGithubClientV4()
}
