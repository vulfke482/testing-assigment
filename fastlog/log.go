package fastlog

// Log is a generic logging interface
type Log interface {
	Info(level uint8, message string, args ...interface{})
	Important(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Trace(message string, args ...interface{})
	Error(message string, args ...interface{})
	WillTrace() bool
	WillDebug() bool
}
