package neo4j

import (
	"github-data-evaluator/config"
	"github-data-evaluator/pkg/utils"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var driver neo4j.Driver
var session neo4j.Session

func init() {
	initNeo4j()
}

// initNeo4j initializes the Neo4j driver.
func initNeo4j() {
	Neo4jConf := config.Conf.Neo4j
	uri := "bolt://" + Neo4jConf.Host + ":" + Neo4jConf.Port
	_driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(Neo4jConf.UserName, Neo4jConf.Password, ""))
	if err != nil {
		utils.GetLogger().Fatal(err.Error())
	}
	driver = _driver
	utils.GetLogger().Info("Neo4j driver initialized")
	session = driver.NewSession(neo4j.SessionConfig{DatabaseName: Neo4jConf.Database})
	utils.GetLogger().Info("Neo4j session initialized")
}

// CloseNeo4j closes the Neo4j driver and session.
func CloseNeo4j() {
	if session != nil {
		session.Close()
	}
	if driver != nil {
		driver.Close()
	}
}
