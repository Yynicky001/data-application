package config

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	Server *Server
	Mysql  *Mysql
}

type Server struct {
	Domain  string `yaml:"domain"`
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
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
}

func InitConfig() {
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
