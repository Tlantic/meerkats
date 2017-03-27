package meerkats

import "time"

type DoneCallback func()

func NopCallback() {}

var _ DoneCallback = NopCallback

type Handler interface {
	Encoder
	LoggerOption

	SetLevel(level Level)
	GetLevel() Level

	Log(time.Time, Level, string, []Field, map[string]string, DoneCallback)

	Clone() Handler
	Dispose()
}
