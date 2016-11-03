package web

import (
	"fmt"
	"log/syslog"
)

//Logger syslog
type Logger struct {
	wrt *syslog.Writer
}

//Debug debug
func (p *Logger) Debug(msg string, args ...interface{}) {
	p.wrt.Debug(fmt.Sprintf(msg, args...))
}

//Info info
func (p *Logger) Info(msg string, args ...interface{}) {
	p.wrt.Info(fmt.Sprintf(msg, args...))
}

//Notice notice
func (p *Logger) Notice(msg string, args ...interface{}) {
	p.wrt.Notice(fmt.Sprintf(msg, args...))
}

//Alert alert
func (p *Logger) Alert(msg string, args ...interface{}) {
	p.wrt.Alert(fmt.Sprintf(msg, args...))
}

//Warning warning
func (p *Logger) Warning(msg string, args ...interface{}) {
	p.wrt.Warning(fmt.Sprintf(msg, args...))
}

//Error error
func (p *Logger) Error(msg string, args ...interface{}) {
	p.wrt.Err(fmt.Sprintf(msg, args...))
}

//NewLogger new logger
func NewLogger(priority syslog.Priority, tag string) (*Logger, error) {
	wrt, err := syslog.New(priority, tag)
	return &Logger{wrt: wrt}, err
}
