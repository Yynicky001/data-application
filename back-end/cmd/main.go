package main

import (
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/db"
)

func main() {
	loading()
	utils.FetchStart()

}

func loading() {
	utils.InitLog()
	config.InitConfig()
	db.InitDB()
	utils.InitGithubClient()
}
