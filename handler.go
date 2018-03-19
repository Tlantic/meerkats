package meerkats

import (
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type DoneCallback func()

func NopCallback() {}

var _ DoneCallback = NopCallback

type Handler interface {
	Encoder
	LoggerOption

	SetLevel(level Level)
	GetLevel() Level

	Log(time.Time, Level, string, []log.Field, map[string]interface{}, DoneCallback)

	Child() Handler
	Close() error
}
