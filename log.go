package flygo

import (
	"fmt"
	"time"
)

type log struct {
}

//Info level log
func (log *log) info(message string, args ...interface{}) {
	log.log("INFO", message, args...)
}

//Warn level log
func (log *log) warn(message string, args ...interface{}) {
	log.log("WARN", message, args...)
}

//Fatal level log
func (log *log) fatal(message string, args ...interface{}) {
	log.log("FATAL", message, args...)
}

//log
func (log *log) log(t, message string, args ...interface{}) {
	str := message
	if args != nil {
		str = fmt.Sprintf(message, args...)
	}
	now := time.Now().Format("2006/01/02 15:04:05.000")
	s := fmt.Sprintf("[%s][FLYGO][%s] %s", now, t, str)
	fmt.Println(s)
}

//Log info
func (a *App) LogInfo(message string, args ...interface{}) {
	a.log.info(message, args...)
}

//Log warn
func (a *App) LogWarn(message string, args ...interface{}) {
	a.log.warn(message, args...)
}

//Log fatal
func (a *App) LogFatal(message string, args ...interface{}) {
	a.log.fatal(message, args...)
}
