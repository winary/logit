package logit

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type LogLevel struct {
	Id   int
	Text string
}

var (
	LevelDebug = LogLevel{1, "debug"}
	LevelInfo  = LogLevel{2, "info"}
	LevelWarn  = LogLevel{3, "warn"}
	LevelError = LogLevel{4, "error"}
	LevelFatal = LogLevel{5, "fatal"}
)

type LogWriter interface {
	Writef(level LogLevel, str string)
}

var (
	Level  LogLevel
	Writer LogWriter
)

// ------------------------------------------------
type LogPkg struct {
	pkg string
}

func NewLogPkg(pkg string) (ret *LogPkg) {
	ret = &LogPkg{
		pkg: pkg,
	}

	return
}

func (this *LogPkg) Debug(format string, args ...interface{}) {
	this.do(LevelDebug, format, args...)
}

func (this *LogPkg) Info(format string, args ...interface{}) {
	this.do(LevelInfo, format, args...)
}

func (this *LogPkg) Warn(format string, args ...interface{}) {
	this.do(LevelWarn, format, args...)
}

func (this *LogPkg) Error(format string, args ...interface{}) {
	this.do(LevelError, format, args...)
}

func (this *LogPkg) Fatal(format string, args ...interface{}) {
	this.do(LevelFatal, format, args...)
	os.Exit(-1)
}

// ------------------------------------------------

func (this *LogPkg) do(level LogLevel, format string, args ...interface{}) {
	if level.Id < Level.Id {
		return
	}

	str := fmt.Sprintf(format, args...)
	log := fmt.Sprintf("%s, %s", this.fileLine(3), str)

	Writer.Writef(level, log)
}

func (this *LogPkg) fileLine(skip int) (str string) {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		str = fmt.Sprintf("file=%s/%s:%d", this.pkg, path.Base(file), line)
	}
	return
}
