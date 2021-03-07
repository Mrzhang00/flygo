package log

import (
	"fmt"
	"log"
	"os"
)

// Logger interface
type Logger interface {
	// Info log
	Info(msg string, args ...interface{})
	// Debug log
	Debug(msg string, args ...interface{})
	// Warn log
	Warn(msg string, args ...interface{})
	// Error log
	Error(msg string, args ...interface{})
}

// DefaultLogger logger
var DefaultLogger Logger = New("[Logger]")

// New return new logger
func New(prefix string) Logger {
	return &logger{
		outLogger: log.New(os.Stdout, prefix, log.LstdFlags),
		errLogger: log.New(os.Stderr, prefix, log.LstdFlags),
	}
}

type logger struct {
	outLogger *log.Logger
	errLogger *log.Logger
}

// Info implements
func (l *logger) Info(msg string, args ...interface{}) {
	l.outLogger.Println("[INFO]" + fmt.Sprintf(msg, args...))
}

// Debug implements
func (l *logger) Debug(msg string, args ...interface{}) {
	l.outLogger.Println("[DEBUG]" + fmt.Sprintf(msg, args...))
}

// Warn implements
func (l *logger) Warn(msg string, args ...interface{}) {
	l.outLogger.Println("[WARN]" + fmt.Sprintf(msg, args...))
}

// Error implements
func (l *logger) Error(msg string, args ...interface{}) {
	l.errLogger.Println("[ERROR]" + fmt.Sprintf(msg, args...))
}
