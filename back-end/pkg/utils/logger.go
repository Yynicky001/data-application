package utils

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var logger *Logger

// Logger 封装后的日志结构
type Logger struct {
	logger     *logrus.Logger
	fileLogger *logrus.Logger
}

// 初始化函数
func init() {
	logrusInit()
}

func logrusInit() {
	// 创建控制台 logger
	consoleLogger := logrus.New()
	consoleLogger.SetFormatter(generateFormatter())
	consoleLogger.SetOutput(colorable.NewColorableStdout())

	// 创建文件 logger
	fileLogger := logrus.New()
	fileFormatter := generateFormatter()
	// 设置文件输出格式
	fileFormatter.DisableColors = true
	fileFormatter.ForceColors = false
	fileLogger.SetFormatter(fileFormatter)
	// 获取文件输出路径
	src, err := setOutputFile()
	if err != nil {
		panic(err)
	}
	// 将文件输出到文件
	fileLogger.SetOutput(src)

	// 创建文件Hook
	fileHook := NewFileHook(src, fileFormatter)

	// 将文件Hook添加到控制台 logger 中
	consoleLogger.AddHook(fileHook)

	// 将控制台 logger 赋值给全局变量
	logger = &Logger{
		logger:     consoleLogger,
		fileLogger: fileLogger,
	}
}

func generateFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              true,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
}

// GetLogger 获取封装后的文本Logger实例
func GetLogger() *Logger { return logger }

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

// WithFields 为 Logger 添加结构化字段并返回新的 Logger
// 简化的WithFields方法，使用变长参数
func (l *Logger) WithFields(fields ...interface{}) *Logger {
	if len(fields)%2 != 0 {
		l.logger.Warn("Invalid number of arguments for WithFields")
		return l
	}

	fieldMap := logrus.Fields{}
	for i := 0; i < len(fields); i += 2 {
		key, okKey := fields[i].(string)
		if !okKey {
			l.logger.Warn("Key must be a string")
			continue
		}
		fieldMap[key] = fields[i+1]
	}

	return &Logger{
		logger: l.logger.WithFields(fieldMap).Logger,
	}
}

// Info 级别的日志记录方法
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

// Infof Info级别的格式化日志记录方法
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn 级别的日志记录方法
func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Error 级别的日志记录方法
func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

// Errorf Error级别的格式化日志记录方法
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal 级别的日志记录方法
func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

// FileHook 用于将日志同时输出到文件
type FileHook struct {
	file      *os.File
	formatter logrus.Formatter
}

func NewFileHook(file *os.File, formatter logrus.Formatter) *FileHook {
	return &FileHook{file: file, formatter: formatter}
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	line, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.file.Write(line)
	return err
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
