package log

import (
	"fmt"
	"log"
	"os"
)

//Define Logger interface
type Logger interface {
	Info(msg string, args ...interface{})  //Info level
	Debug(msg string, args ...interface{}) //Debug level
	Warn(msg string, args ...interface{})  //Warn level
	Error(msg string, args ...interface{}) //Error level
}

var DefaultLogger Logger = New("[Logger]")

func New(prefix string) Logger {
	return &logger{
		olog: log.New(os.Stdout, prefix, log.LstdFlags),
		elog: log.New(os.Stderr, prefix, log.LstdFlags),
	}
}

type logger struct {
	olog *log.Logger
	elog *log.Logger
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.olog.Println("[INFO]" + fmt.Sprintf(msg, args...))
}

func (l *logger) Debug(msg string, args ...interface{}) {
	l.olog.Println("[DEBUG]" + fmt.Sprintf(msg, args...))
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.olog.Println("[WARN]" + fmt.Sprintf(msg, args...))
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.elog.Println("[ERROR]" + fmt.Sprintf(msg, args...))
}
