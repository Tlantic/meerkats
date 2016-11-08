package tests

import (
	"testing"
	. "github.com/Tlantic/meerkats"
)
/*
func TestNewMeerkat(t *testing.T) {
	out := handlers.NewWritterLogger( os.Stdout )

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:10,
		TimeLayout: time.RFC3339Nano,
	})
	defer sentry.Close()


	sentry.RegisterHandler(LEVEL_ALL, out.HandleEntry)

	sentry.Log(LEVEL_INFO, "Test")
}*/


func BenchmarkMeerkatLog_W1(b *testing.B) {

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:1,
		QueueSize: 1000,
	})

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sentry.Trace("Test")
	}
}

func BenchmarkMeerkatLog_W5(b *testing.B) {

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:5,
		QueueSize: 1000,
	})

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sentry.Trace("Test")
	}
}

func BenchmarkMeerkatLog_W100(b *testing.B) {

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:100,
		QueueSize: 1000,
	})

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sentry.Trace("Test")
	}
}

func BenchmarkMeerkatLog_W250(b *testing.B) {

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:250,
		QueueSize: 1000,
	})

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sentry.Trace("Test")
	}
}

func BenchmarkMeerkatLog_W1000(b *testing.B) {

	sentry := NewMeerkat(MeerkatOptions{
		MaxWorkers:1000,
		QueueSize: 1000,
	})

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		sentry.Trace("Test")
	}
}