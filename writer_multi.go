package logit

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

// ------------------------------------------------

type MultiWriter struct {
	filename string
	ts       time.Time
	mutex    sync.Mutex
	filer    *os.File
	writer   io.Writer
}

func UseMultiWriter(level LogLevel, filename string) (err error) {
	if "" == filename {
		err = fmt.Errorf("filename is empty")
		return
	}
	err = os.MkdirAll(path.Dir(filename), 0644)
	if nil != err {
		return
	}

	Level, Writer = level, &MultiWriter{
		filename: filename,
		ts:       time.Now(),
	}

	return
}

func (this *MultiWriter) newWriter(ts time.Time) (err error) {
	if nil != this.filer {
		this.filer.Close()
	}

	var (
		name = fmt.Sprintf("%s.%s.log", this.filename, ts.Format("2006-01-02"))
	)

	this.filer, err = os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if nil != err {
		return
	}
	this.writer = io.MultiWriter(os.Stdout, this.filer)

	return
}

func (this *MultiWriter) setWriter(ts time.Time) (err error) {
	need_new := func() bool {
		if nil == this.writer {
			return true
		}
		if this.ts.Year() != ts.Year() || this.ts.Month() != ts.Month() || this.ts.Day() != ts.Day() {
			return true
		}
		return false
	}
	todo_new := func() error {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		if need_new() {
			this.ts = ts
			return this.newWriter(ts)
		}

		return nil
	}

	if need_new() {
		return todo_new()
	}

	return
}

func (this *MultiWriter) Writef(level LogLevel, log_str string) {
	ts := time.Now()

	err := this.setWriter(ts)
	if nil != err {
		log.Printf("Writef err %v", err)
		return
	}

	// step 1
	buf := &bytes.Buffer{}
	str := fmt.Sprintf("time=%s.%03d, level=%s, ",
		ts.Format("2006-01-02 15:04:05"), ts.UnixNano()/1e6%1e3, level.Text)
	buf.WriteString(str)

	// step 2
	buf.WriteString(log_str + "\n")

	// step 3
	this.writer.Write(buf.Bytes())

	return
}
