package handlers

import (
	"os"
	"sync"
	. "github.com/Tlantic/meerkats"
	"io"
	"fmt"
)

//noinspection GoUnusedConst
const (
	COLOR_NONE = 0
	COLOR_RED = 31
	COLOR_GREEN  = 32
	COLOR_YELLOW = 33
	COLOR_BLUE = 34
	COLOR_GRAY = 37
)

var COLORS = [...]int{
	LEVEL_TRACE: COLOR_NONE,
	LEVEL_DEBUG: COLOR_GRAY,
	LEVEL_INFO:  COLOR_BLUE,
	LEVEL_WARNING:  COLOR_YELLOW,
	LEVEL_ERROR: COLOR_RED,
	LEVEL_FATAL: COLOR_RED,
	LEVEL_PANIC: COLOR_RED,
}

type WriterLogger struct {
	w	io.Writer
	mu	sync.Mutex
}
func NewWritterLogger(w io.Writer) *WriterLogger {
	return &WriterLogger{
		w: w,
		mu: sync.Mutex{},
	}
}

func (l *WriterLogger) HandleEntry(e Entry, done Callback) {
	defer done()

	c := COLORS[e.Level]

	text := fmt.Sprintf("\033[%dm timestamp=\"%s\" level=\"%s\" message=\"%s\" %s\033[0m\n",  c, e.Timestamp.Format(e.TimeLayout), e.Level, e.Message, e.Fields.String())

	switch e.Level {
	case LEVEL_FATAL:
		fmt.Fprintf(l.w, text)
		os.Exit(1)
	case LEVEL_PANIC:
		panic(text)
	default:
		fmt.Fprintf(l.w, text)
	}

}