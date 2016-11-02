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


	for i := 0; i < 100000; i++ {
		go sentry.WithField("hello", "world").WithFields(Fields{
			"remaining": 100000 - i,
		}).Warning(i)
	}

	sentry.WithField("hello", "world").WithFields(Fields{
		"make": "better",
	}).Error("Opps something went wrong")
}