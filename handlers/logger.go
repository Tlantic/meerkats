package handlers

import (
	"log"
	"sync"
	. "github.com/Tlantic/meerkats"
)

type Logger struct {
	mu	sync.Mutex
}
func NewLogger() *Logger {
	return &Logger{
		mu: sync.Mutex{},
	}
}

func (l *Logger) HandleEntry(e Entry) {

	switch e.Level {
	case LEVEL_FATAL:
		log.Fatalln(e.String())
	case LEVEL_PANIC:
		log.Panicln(e.String())
	default:
		log.Println(e.String())
	}
}