package neo4j

import (
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var neo4jDriver neo4j.Driver

func init() {
	initNeo4j()
}

// initNeo4j initializes the Neo4j driver.
func initNeo4j() {
	Neo4jConf := config.Conf.Neo4j
	uri := "bolt://" + Neo4jConf.Host + ":" + Neo4jConf.Port
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(Neo4jConf.UserName, Neo4jConf.Password, ""))
	if err != nil {
		utils.GetLogger().Fatal(err.Error())
	}
	neo4jDriver = driver
}

// CloseNeo4j closes the Neo4j driver connection.
func CloseNeo4j() {
	if neo4jDriver != nil {
		neo4jDriver.Close()
	}
}
