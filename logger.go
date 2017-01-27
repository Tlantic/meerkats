package meerkats

import (
	"fmt"
)

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

	Write(p []byte) (n int, err error)
	Clone() Logger
	Dispose()
}


type StandardLogger struct {
	Logger
	level  Level
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