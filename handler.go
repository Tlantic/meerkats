package meerkats

import "time"




type HandlerSet []Handler


type Handler interface {
	Encoder

	SetLevel(level Level)
	GetLevel() Level

	Log(time.Time, Level, string, []Field, map[string]string)

	Clone() Handler
	Dispose()
}
