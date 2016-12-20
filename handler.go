package meerkats



type HandlerOption func(Handler)

func TimeLayoutOption(layout string) HandlerOption {
	return HandlerOption(func(handler Handler) {
		handler.SetTimeLayout(layout)
	})
}

func LevelOption(level Level) HandlerOption {
	return HandlerOption(func(handler Handler) {
		handler.SetLevel(level)
	})
}


type HandlerSet []Handler


type Handler interface {
	Encoder

	Log(Level, string, []KeyValue)

	SetTimeLayout(layout string)
	GetTimeLayout() string

	SetLevel(level Level)
	GetLevel() Level


	Clone() Handler
	Dispose()
}


