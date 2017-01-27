package meerkats

import (
	"testing"
	"github.com/Tlantic/meerkats/handlers/writer"
)

func TestNew(t *testing.T) {
	logger := New(TRACE)
	if (logger == nil) {
		t.Fail()
	}
}


func TestRegister(t *testing.T) {
	logger := New(TRACE)
	logger.Register(writer.New(LevelOption(LEVEL_ALL)))
}


func TestLog(t *testing.T) {
	logger := New(TRACE)
	logger.Log(TRACE, "test message", Field{ Key: "boolValue", Type: TypeBool, ValueBool: true })
}