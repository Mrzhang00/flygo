package flygo

import (
	"fmt"
	"log"
	"os"
)

//Info level log
func (a *App) Info(message string, args ...interface{}) {
	a.outLogger.Printf("%s", fmt.Sprintf(message, args...))
}

//Warn level log
func (a *App) Warn(message string, args ...interface{}) {
	a.outLogger.Printf("%s", fmt.Sprintf(message, args...))
}

//Error level log
func (a *App) Error(message string, args ...interface{}) {
	a.errLogger.Printf("%s", fmt.Sprintf(message, args...))
}

func (a *App) setLoggers() {
	var outLogger, errLogger *log.Logger
	if app.Config.Flygo.Log.Type == "stdout" {
		outLogger = log.New(os.Stdout, "[FLYGO]", log.Flags())
		errLogger = log.New(os.Stderr, "[FLYGO]", log.Flags())
	} else if app.Config.Flygo.Log.Type == "file" {
		outFile, ofe := os.Open(app.Config.Flygo.Log.File.Out)
		if ofe != nil || outFile == nil {
			outFile, _ = os.Create(app.Config.Flygo.Log.File.Out)
		}
		errFile, efe := os.Open(app.Config.Flygo.Log.File.Err)
		if efe != nil || errFile == nil {
			errFile, _ = os.Create(app.Config.Flygo.Log.File.Err)
		}
		outLogger = log.New(outFile, "[FLYGO]", log.LstdFlags)
		errLogger = log.New(errFile, "[FLYGO]", log.LstdFlags)
	}
	a.outLogger = outLogger
	a.errLogger = errLogger
}
