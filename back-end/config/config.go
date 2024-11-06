package config

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
	GitHub *GitHub `yaml:"github"`
	Neo4j  *Neo4j  `yaml:"neo4j"`
}

type Server struct {
	Domain  string `yaml:"domain"`
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
	Command string `yaml:"command"`
}

type Mysql struct {
	DriverName        string `yaml:"driverName"`
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	UserName          string `yaml:"username"`
	Password          string `yaml:"password"`
	Database          string `yaml:"database"`
	Charset           string `yaml:"charset"`
	MaxIdleConn       int    `yaml:"maxIdleConn"`
	MaxOpenConn       int    `yaml:"maxOpenConn"`
	MaxLifetime       int    `yaml:"maxLifetime"`
	DefaultStringSize uint   `yaml:"defaultStringSize"`
	Migrate           bool   `yaml:"migrate"`
}

type GitHub struct {
	Token     []string `yaml:"token"`
	Repo      bool     `yaml:"repo"`
	Developer bool     `yaml:"developer"`
	Strategy  string   `yaml:"github_api_strategy"`
}

type Neo4j struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
