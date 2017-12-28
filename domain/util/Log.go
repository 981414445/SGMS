package util

import (
	"log"
	"os"
)

const (
	LOG_TRACE = iota
	LOG_DEBUG
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
)

type Log struct {
	logFile *os.File
	log     *log.Logger
	level   int
}

func NewLog(file, flag string, level int) *Log {
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}
	r := new(Log)
	r.logFile = logFile
	r.log = log.New(logFile, flag, log.LstdFlags)
	r.level = level
	return r
}

func (this *Log) Trace(v ...interface{}) {
	if this.level > LOG_TRACE {
		return
	}
	this.log.SetPrefix("[INFO]")
	this.log.Println(v...)
}
func (this *Log) Debug(v ...interface{}) {
	if this.level > LOG_DEBUG {
		return
	}
	this.log.SetPrefix("[INFO]")
	this.log.Println(v...)
}
func (this *Log) Info(v ...interface{}) {
	if this.level > LOG_INFO {
		return
	}
	this.log.SetPrefix("[INFO]")
	this.log.Println(v...)
}

func (this *Log) Warn(v ...interface{}) {
	if this.level > LOG_WARN {
		return
	}
	this.log.SetPrefix("[WARN]")
	this.log.Println(v...)
}

func (this *Log) Error(v ...interface{}) {
	if this.level > LOG_ERROR {
		return
	}
	this.log.SetPrefix("[ERROR]")
	this.log.Println(v...)
}
