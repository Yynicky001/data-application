package config

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
	Neo4j  *Neo4j  `yaml:"neo4j"`
	ES     *ES     `yaml:"es"`
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

type Neo4j struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ES struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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
