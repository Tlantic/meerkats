package meerkats

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type LoggerReceiver func(Logger)

func (f LoggerReceiver) Apply(l Logger) {
	f(l)
}

type LoggerOption interface {
	Apply(Logger)
}

func WithSpan(span opentracing.Span) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		l.WithSpan(span)
	})
}

func WithLevel(level Level) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		l.SetLevel(level)
	})
}

func WithFields(fs ...log.Field) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		l.EmitField(fs...)
	})
}

func WithTag(meta map[string]string) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		for k, v := range meta {
			l.SetTag(k, v)
		}
	})
}

type HandlerOption interface {
	Apply(Handler)
}

type HandlerReceiver func(Handler)

func (f HandlerReceiver) Apply(h Handler) {
	f(h)
}

func LevelOption(level Level) HandlerOption {
	return HandlerReceiver(func(handler Handler) {
		handler.SetLevel(level)
	})
}
