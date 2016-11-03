package web

import (
	"fmt"
	"log/syslog"
)

//Logger web logger
type Logger struct {
	Writer *syslog.Writer `inject:"logger.web"`
}

//Debug debug
func (p *Logger) Debug(msg string, args ...interface{}) {
	p.Writer.Debug(fmt.Sprintf(msg, args...))
}

//Info info
func (p *Logger) Info(msg string, args ...interface{}) {
	p.Writer.Info(fmt.Sprintf(msg, args...))
}

//Notice notice
func (p *Logger) Notice(msg string, args ...interface{}) {
	p.Writer.Notice(fmt.Sprintf(msg, args...))
}

//Alert alert
func (p *Logger) Alert(msg string, args ...interface{}) {
	p.Writer.Alert(fmt.Sprintf(msg, args...))
}

//Warning warning
func (p *Logger) Warning(msg string, args ...interface{}) {
	p.Writer.Warning(fmt.Sprintf(msg, args...))
}

//Error error
func (p *Logger) Error(msg string, args ...interface{}) {
	p.Writer.Err(fmt.Sprintf(msg, args...))
}

// -----------------------------------------------------------------------------

//JobLogger job logger
type JobLogger struct {
	Writer *syslog.Writer `inject:"logger.jobs"`
}

//Print print
func (p *JobLogger) Print(args ...interface{}) {
	p.Printf("%+v", args)
}

//Printf printf
func (p *JobLogger) Printf(format string, args ...interface{}) {
	p.Writer.Info(fmt.Sprintf(format, args...))
}

//Println println
func (p *JobLogger) Println(args ...interface{}) {
	p.Print(args...)
}

//Fatal fatal
func (p *JobLogger) Fatal(args ...interface{}) {
	p.Fatalf("%+v", args)
}

//Fatalf fatalf
func (p *JobLogger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	p.Writer.Emerg(msg)
	panic(msg)
}

//Fatalln fatalln
func (p *JobLogger) Fatalln(args ...interface{}) {
	p.Fatal(args...)
}

//Panic panic
func (p *JobLogger) Panic(args ...interface{}) {
	p.Fatal(args...)
}

//Panicf panicf
func (p *JobLogger) Panicf(format string, args ...interface{}) {
	p.Fatalf(format, args...)
}

//Panicln pacnicln
func (p *JobLogger) Panicln(args ...interface{}) {
	p.Fatal(args...)
}
