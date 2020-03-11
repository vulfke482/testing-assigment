package fastlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// These are the log levels supported by WrappedLogger out of the box
// The application can choose to define and use their own levels
// WrappedLogger doesnt make any assumptions and simply uses the level as a measure of message's importance
const (
	LevelImportant = uint8(0)
	LevelDebug     = uint8(1)
	LevelTrace     = uint8(2)
)

// WrappedLogger is a logger that is wrapped with a config file
// It takes a config to determine which level its running at
type WrappedLogger struct {
	*zap.SugaredLogger
	level uint8
}

// NewWrappedLogger creates a new logger which wraps the given config
func NewWrappedLogger(name, outfile string, level uint8) (*WrappedLogger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// By default zapConfig.ErrorOutputPaths will go stderr
	// We should just set the place where we want to put the debug log
	zapConfig.OutputPaths = []string{outfile}
	zapConfig.DisableCaller = true

	l, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	l = l.Named(name)
	return &WrappedLogger{
		SugaredLogger: l.Sugar(),
		level:         level,
	}, nil
}

// Finish is the function to be called before exiting
func (wl WrappedLogger) Finish() error {
	return wl.Sync()
}

// Info prints the given message and data
func (wl WrappedLogger) Info(level uint8, message string, args ...interface{}) {
	if !wl.isLevelEnabled(level) {
		return
	}
	wl.Infow(message, args...)
}

func (wl WrappedLogger) isLevelEnabled(level uint8) bool {
	return level < wl.level
}

// Important is to print important information
func (wl WrappedLogger) Important(message string, args ...interface{}) {
	wl.Info(LevelImportant, message, args...)
}

// Debug prints messages that should be used when debugging
func (wl WrappedLogger) Debug(message string, args ...interface{}) {
	wl.Info(LevelDebug, message, args...)
}

// Trace is to print information for a detailed trace
func (wl WrappedLogger) Trace(message string, args ...interface{}) {
	wl.Info(LevelTrace, message, args...)
}

// Error prints an error message to stderr
func (wl WrappedLogger) Error(message string, args ...interface{}) {
	wl.Errorw(message, args...)
}

// WillTrace returns true if logger is enabled for tracing
func (wl WrappedLogger) WillTrace() bool {
	return wl.isLevelEnabled(LevelTrace)
}

// WillDebug returns true if logger is enabled for debugging
func (wl WrappedLogger) WillDebug() bool {
	return wl.isLevelEnabled(LevelDebug)
}
