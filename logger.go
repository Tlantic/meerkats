package meerkats

type Logger interface {
	Encoder

	Log(level Level, msg string, fields ...KeyValue)
	Trace(msg string, fields ...KeyValue)
	Debug(msg string, fields ...KeyValue)
	Info(msg string, fields ...KeyValue)
	Warn(msg string, fields ...KeyValue)
	Error(msg string, fields ...KeyValue)
	Panic(msg string, fields ...KeyValue)
	Fatal(msg string, fields ...KeyValue)
}