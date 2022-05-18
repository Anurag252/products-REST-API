package logger

import "fmt"

type Logger interface {
	LogError(string)
	LogWarn(string)
	LogMessage(string)
}

type ConsoleLogger struct {
}

func New() Logger {
	logger := ConsoleLogger{}
	return logger
}

func (logger ConsoleLogger) LogError(message string) {
	fmt.Println(fmt.Sprintf("ERROR:- %s", message))
}

func (logger ConsoleLogger) LogWarn(message string) {
	fmt.Println(fmt.Sprintf("WARN:- %s", message))
}

func (logger ConsoleLogger) LogMessage(message string) {
	fmt.Println(fmt.Sprintf("MESSAGE:- %s", message))
}
