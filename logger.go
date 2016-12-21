package meerkats

type Logger interface {
	Encoder

	SetLevel(Level)

	Register(...Handler)

	SetMeta(string, string)
	GetMeta(string) string

	Log(level Level, msg string, fields ...Field)
	Trace(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	Clone() Logger
	Dispose()
}