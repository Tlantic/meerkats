package meerkats

type Level uint8

const (
	LevelTrace   Level = 1 << iota // 00000001 -> 1
	LevelDebug                     // 00000010 -> 2
	LevelInfo                      // 00000100 -> 4
	LevelWarning                   // 00001000 -> 8
	LevelError                     // 00010000 -> 16
	LevelFatal                     // 00100000 -> 32
	LevelPanic                     // 01000000 -> 64
)

const LevelAll = LevelTrace | LevelDebug | LevelInfo | LevelWarning | LevelError | LevelFatal | LevelPanic // -> 01111111 -> 127

var Levels = [...]string{
	LevelTrace:   "trace",
	LevelDebug:   "debug",
	LevelInfo:    "info",
	LevelWarning: "warning",
	LevelError:   "error",
	LevelFatal:   "fatal",
	LevelPanic:   "panic",
}

func (lvl Level) String() string {
	return Levels[lvl]
}

func (lvl Level) Apply(l Logger) {
	l.SetLevel(lvl)
}
