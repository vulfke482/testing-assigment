package fastlog

import (
	"fmt"
	"os"
)

var globalLoggerStdout *WrappedLogger
var globalLoggerStderr *WrappedLogger

// EnvDebugLevel variable for control debug level
const EnvDebugLevel = "DEBUG_LEVEL"

// EnvLogName variable for control name in logs
const EnvLogName = "LOG_NAME"

func init() {
	var err error

	level := LevelDebug
	if val, exists := os.LookupEnv(EnvDebugLevel); exists {
		switch val {
		case "0":
			// This will disable logging
			level = LevelImportant
		case "1":
			// Only LevelImportant will be let through
			level = LevelDebug
		default:
			// Debug messages will go through
			// We dont support tracing at global level
			level = LevelTrace
		}
	}

	name := os.Getenv(EnvLogName)

	globalLoggerStdout, err = NewWrappedLogger(name, "stdout", level)
	if err != nil {
		panic(fmt.Sprintf("Unable to create global stdout logger: %s", err))
	}

	globalLoggerStderr, err = NewWrappedLogger(name, "stderr", LevelDebug)
	if err != nil {
		panic(fmt.Sprintf("Unable to create global stderr logger: %s", err))
	}
}

// Important logs a message to stdinfo
func Important(message string, args ...interface{}) {
	globalLoggerStdout.Important(message, args...)
}

// Debug prints a debug message
func Debug(message string, args ...interface{}) {
	globalLoggerStdout.Debug(message, args...)
}

// Error logs a message to stderr
func Error(message string, args ...interface{}) {
	globalLoggerStderr.Error(message, args...)
}

// Fatal logs a message to stderr and then calls an os.Exit
func Fatal(message string, args ...interface{}) {
	Error(message, args...)
	os.Exit(1)
}
