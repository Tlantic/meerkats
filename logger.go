package meerkats

import (
	. "context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func ContextWithLogger(ctx Context, logger Logger) Context {
	return WithValue(ctx, uniqueKey, logger)
}
func LoggerFromContext(ctx Context) (logger Logger) {
	var ok bool
	if logger, ok = ctx.Value(uniqueKey).(Logger); !ok {
		logger = Clone()
		if span := opentracing.SpanFromContext(ctx); span != nil {
			logger.WithSpan(span)
		}
	}
	return
}

type Logger interface {
	Encoder

	OperationName() string
	SetOperationName(string)

	Span() opentracing.Span
	WithSpan(opentracing.Span)

	SetLevel(Level)

	Register(...Handler)

	// Deprecated: Use SetTag
	SetMeta(string, string)
	SetTag(string, interface{})
	// Deprecated: Use GetTag
	GetMeta(string) string
	GetTag(string) interface{}

	Log(level Level, msg string, fields ...log.Field)
	Trace(msg string, fields ...log.Field)
	Debug(msg string, fields ...log.Field)
	Info(msg string, fields ...log.Field)
	Warn(msg string, fields ...log.Field)
	Error(msg string, fields ...log.Field)
	Panic(msg string, fields ...log.Field)
	Fatal(msg string, fields ...log.Field)

	Write(p []byte) (n int, err error)
	Clone() Logger
	Dispose()
}

type StandardLogger struct {
	Logger
	level Level
}

func NewStandardLogger(logger Logger, printLevel Level) *StandardLogger {
	return &StandardLogger{logger, printLevel}
}

func (l *StandardLogger) Printf(format string, v ...interface{}) {
	l.Log(l.level, fmt.Sprintf(format, v...))
}

// Print calls l.Log to print to the logger at the specified level.
// Arguments are handled in the manner of fmt.Print.
func (l *StandardLogger) Print(v ...interface{}) {
	l.Log(l.level, fmt.Sprint(v...))
}

// Println calls l.Log to print to the logger at the specified level.
// Arguments are handled in the manner of fmt.Println.
func (l *StandardLogger) Println(v ...interface{}) {
	l.Log(l.level, fmt.Sprintln(v...))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *StandardLogger) Fatal(v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintln(v...))
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *StandardLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, v...))
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (l *StandardLogger) Fatalln(v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintln(v...))
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *StandardLogger) Panic(v ...interface{}) {
	l.Logger.Panic(fmt.Sprintln(v...))
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *StandardLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panic(fmt.Sprintf(format, v...))
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *StandardLogger) Panicln(v ...interface{}) {
	l.Logger.Panic(fmt.Sprintln(v...))
}

type JaegerLogger struct {
	Logger
}

// Infof logs a message at info priority
func (l *JaegerLogger) Error(msg string) {
	l.Logger.Error(msg)
}
func (l *JaegerLogger) Infof(msg string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(msg, args...))
}
