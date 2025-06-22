package log

import (
	"fmt"

	"go.uber.org/zap"
)

// env: local -> std output
// env: prod -> file
const (
	debugLevel = 1
	infoLevel  = 2
	errorLevel = 3
)

var GlobalLog *Log

type Log struct {
	zap   *zap.Logger
	level int
}

type LogField struct {
	Key   string
	Value interface{}
}

type Config struct {
	LogLevel int
	FilePath string
}

// Init initialize global Log using configs
func Init(conf Config) error {
	// TODO: hardcode local
	GlobalLog = &Log{
		zap:   nil,
		level: debugLevel,
	}
	return nil
}

func Debug(msg string, field ...LogField) {
	if GlobalLog.level <= debugLevel {
		m := make(map[string]interface{})
		fmt.Println(msg, m)
	}
}

func Info(msg string, field ...LogField) {
	if GlobalLog.level <= infoLevel {
		m := make(map[string]interface{})
		fmt.Println(msg, m)
	}
}

func Error(msg string, field ...LogField) {
	if GlobalLog.level <= errorLevel {
		m := make(map[string]interface{})
		fmt.Println(msg, m)
	}
}
