package es

import (
	"github-data-evaluator/config"
	"github.com/olivere/elastic/v7"
)

var _client *elastic.Client

func init() {
	client, err := elastic.NewClient(
		elastic.SetURL(config.Conf.ES.Host+":"+config.Conf.ES.Port), // 设置ES地址
		elastic.SetSniff(false), // 禁用节点嗅探
	)
	if err != nil {
		panic(err)
	}
	_client = client
}
