package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		return
	}
	src, err := setOutputFile()
	if err != nil {
		panic(err)
	}
	//日志对象实例化
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)         //设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{ //设置日志格式
		TimestampFormat: time.DateTime,
	})
	LogrusObj = logger
}

// 设置日志输出文件
func setOutputFile() (*os.File, error) {
	now := time.Now()                       //获取当前时间
	logFilePath := ""                       //日志文件路径
	if dir, err := os.Getwd(); err != nil { //获取当前目录
		return nil, err
	} else {
		logFilePath = dir + "/logs"
	}
	_, err := os.Stat(logFilePath) //判断是否有logs文件夹
	if os.IsNotExist(err) {        //没有就创建
		if err := os.MkdirAll(logFilePath, 0755); err != nil { //创建文件夹
			return nil, err
		}
	}

	logFileName := now.Format(time.DateOnly) + ".log" //日志文件名
	fileName := filepath.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666) //打开日志文件
	if err != nil {
		return nil, err
	}
	return src, nil
}
