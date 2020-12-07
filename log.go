package flygo

import (
	"fmt"
	l "log"
	"os"
)

//Define log struct
type log struct {
	ologger *l.Logger
	elogger *l.Logger
}

//Info level log
func (l *log) Info(message string, args ...interface{}) {
	l.ologger.Printf("%s", fmt.Sprintf(message, args...))
}

//Warn level log
func (l *log) Warn(message string, args ...interface{}) {
	l.ologger.Printf("%s", fmt.Sprintf(message, args...))
}

//Error level log
func (l *log) Error(message string, args ...interface{}) {
	l.elogger.Printf("%s", fmt.Sprintf(message, args...))
}

func (a *App) setLoggers() {
	var of, ef = os.Stdout, os.Stderr
	if a.Config.Flygo.Log.Type == "file" {
		outFile, ofe := os.Open(a.Config.Flygo.Log.File.Out)
		if ofe != nil || outFile == nil {
			outFile, _ = os.Create(a.Config.Flygo.Log.File.Out)
		}
		errFile, efe := os.Open(a.Config.Flygo.Log.File.Err)
		if efe != nil || errFile == nil {
			errFile, _ = os.Create(a.Config.Flygo.Log.File.Err)
		}
	}
	a.Logger = &log{
		ologger: l.New(of, fmt.Sprintf("[%s]", a.Name), l.LstdFlags),
		elogger: l.New(ef, fmt.Sprintf("[%s]", a.Name), l.LstdFlags),
	}
}
