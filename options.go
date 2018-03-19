package meerkats

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)


// LoggerOption represents an object that can act as a logger visitor
type LoggerOption interface {
	Apply(Logger)
}

// LoggerOptionFunc represents a function that can act as a logger visitor
type LoggerOptionFunc func(Logger)
func (f LoggerOptionFunc) Apply(l Logger) {
	f(l)
}

// WithSpan makes use of an opentracing.Span. Useful when logging metrics
// and traces to remote providers like Jaeger.
func WithSpan(span opentracing.Span) LoggerOption {
	return LoggerOptionFunc(func(l Logger) {
		l.WithSpan(span)
	})
}

// WithLevel sets the lowest level that can cause the logger the trigger its handlers.
func WithLevel(level Level) LoggerOption {
	return LoggerOptionFunc(func(l Logger) {
		l.SetLevel(level)
	})
}
// WithFields adds fields to the logger and consequent children.
func WithFields(fs ...log.Field) LoggerOption {
	return LoggerOptionFunc(func(l Logger) {
		l.EmitField(fs...)
	})
}
// WithTag adds a tag to the logger and consequent children.
func WithTag(k string, v interface{}) LoggerOption {
	return LoggerOptionFunc(func(l Logger) {
			l.SetTag(k, v)
	})
}
// WithTags adds tags to the logger and consequent children.
func WithTags(tags map[string]interface{}) LoggerOption {
	return LoggerOptionFunc(func(l Logger) {
		for k, v := range tags {
			l.SetTag(k, v)
		}
	})
}

// HandlerOption represents an object that can act as a handler visitor
type HandlerOption interface {
	Apply(Handler)
}

// HandlerOptionFunc represents a function that can act as a handler visitor
type HandlerOptionFunc func(Handler)
func (f HandlerOptionFunc) Apply(h Handler) {
	f(h)
}

// LevelOption sets the levels the the handler will receive logs,
// unlike the logger itself were the level determines the lowest level that he will consider,
// to the handlers this option establishes at which levels the handler will operate on.
//
// Examples:
// 		All Levels -> LevelTrace | LevelDebug | LevelInfo | LevelWarning | LevelError | LevelFatal | LevelPanic
//		Only at Info, Warning and Error -> LevelInfo | LevelWarning | LevelError
func LevelOption(level Level) HandlerOption {
	return HandlerOptionFunc(func(handler Handler) {
		handler.SetLevel(level)
	})
}
