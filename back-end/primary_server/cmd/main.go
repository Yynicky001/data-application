package main

import (
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
	utils.GetLogger().Info("start fetch data")
	// TODO: fetch data
}
