package zlog

import (
	"bytes"
	"context"
	"io"
	"os"
	"time"
)

type ZLogger struct {
	handler Handler
}

func NewZLogger(handler Handler) *ZLogger {
	return &ZLogger{handler: handler}
}

func Default() *ZLogger {
	return NewZLogger(DefaultHandler())
}

type Handler interface {
	Handle(context.Context, Record) error
}

type Record struct {
	// The time at which the output method (Log, Info, etc.) was called.
	Time time.Time

	// The log message.
	Message string

	// The level of the event.
	Level Level
}

type handleState struct {
	handler *Handler
	sep     string
}

type defaultHandler struct {
	out io.Writer
}

func (d defaultHandler) Handle(ctx context.Context, record Record) error {
	buf := bytes.Buffer{}
	buf.WriteString(time.Now().String())
	buf.WriteByte(' ')
	level := "INFO "
	switch record.Level {
	case LevelDebug:
		level = "DEBUG "
	case LevelInfo:
	case LevelWarn:
		level = "WARN "
	case LevelError:
		level = "ERROR "
	}
	buf.WriteString(level)
	buf.WriteString(record.Message)
	buf.WriteByte('\n')
	_, err := d.out.Write(buf.Bytes())
	return err
}

func newDefaultHandler(out io.Writer) *defaultHandler {
	return &defaultHandler{out: out}
}

func DefaultHandler() Handler {
	return newDefaultHandler(os.Stdout)
}

type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

func (l *ZLogger) Info(msg string) error {
	return l.handler.Handle(context.Background(), Record{
		Time:    time.Now(),
		Message: msg,
		Level:   0,
	})
}

func (l *ZLogger) Debug(msg string) error {
	return l.handler.Handle(context.Background(), Record{
		Time:    time.Now(),
		Message: msg,
		Level:   -4,
	})
}
