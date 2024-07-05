package log

import (
	"fmt"
	"log"
	"os"
)

type LogLevel string

var (
	InfoLevel    LogLevel = "INFO"
	WarningLevel LogLevel = "WARNING"
	ErrorLevel   LogLevel = "ERROR"
	FatalLevel   LogLevel = "FATAL"
)

const TemplateLogMessage = "[%s] %s\n"

func Info(message string, args ...any) {
	logMessage(InfoLevel, message, args)
}

func Warning(message string, args ...any) {
	logMessage(WarningLevel, message, args)
}

func Error(message string, args ...any) {
	logMessage(ErrorLevel, message, args)
}

func Fatal(message string, args ...any) {
	logMessage(FatalLevel, message, args)
	os.Exit(1)
}

func logMessage(level LogLevel, message string, args ...any) {
	message = fmt.Sprintf(message, args)

	log.Printf(TemplateLogMessage, level, message)
}
