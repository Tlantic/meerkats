package meerkats

import (
	"testing"

	"github.com/Tlantic/meerkats/handlers/writer"
	"github.com/opentracing/opentracing-go/log"
)

func TestNew(t *testing.T) {
	logger := New(LevelTrace)
	if logger == nil {
		t.Fail()
	}
}

func TestRegister(t *testing.T) {
	logger := New(LevelTrace)
	logger.Register(writer.New(LevelOption(LevelAll)))
}

func TestLog(t *testing.T) {
	logger := New(LevelTrace)
	logger.Log(LevelTrace, "test message", log.Bool("boolValue", true))
}
