package meerkats


type ContextOption func(*Context)


func WithLevel(level Level) ContextOption {
	return ContextOption(func(ctx *Context) {
		ctx.level = level
	})
}

func WithFields(fs ...KeyValue) ContextOption {
	return ContextOption(func(ctx *Context) {
		ctx.With(fs...)
	})
}