package meerkats


type LoggerReceiver func(Logger)

func (f LoggerReceiver) Apply(l Logger) {
	f(l)
}

type LoggerOption interface {
	Apply(Logger)
}

func WithLevel(level Level) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		l.SetLevel(level)
	})
}

func WithFields(fs ...Field) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		l.With(fs...)
	})
}

func WithMeta(meta map[string]string) LoggerOption {
	return LoggerReceiver(func(l Logger) {
		for k, v := range meta {
			l.SetMeta(k, v)
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