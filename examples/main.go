package main

import (
	. "github.com/Tlantic/meerkats"
	"github.com/Tlantic/meerkats/handlers"
	"log"
	"os"
	"time"
)

func main() {

	f, err := os.OpenFile("out.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)
	l :=  handlers.NewLogger()


	out := handlers.NewWritterLogger( os.Stdout )

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:10,
		TimeLayout: time.RFC3339Nano,
	})
	defer sentry.Close()


	sentry.RegisterHandler(LEVEL_ALL, out.HandleEntry, l.HandleEntry)


	sentry.With(MeerkatOptions{
		MaxWorkers:10,
		TimeLayout: time.RFC3339Nano,
	}).Error("Opps something went wrong")
}