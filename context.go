package meerkats

import (
	"time"
	"sync"
)

var now = time.Now

type Metadata map[string]string

type Context struct {
	Metadata
	level    Level

	handlers []Handler
	length   int
	mu       sync.Mutex
	wg       sync.WaitGroup
}

func New(level Level, options...ContextOption) (ctx *Context) {
	ctx = &Context{
		handlers: make([]Handler, 0, 0),
		level: level,
		Metadata: map[string]string{},
	}
	for _, opt := range options {
		opt(ctx)
	}
	return
}
func From( parent *Context, options...ContextOption) (ctx *Context) {
	ctx = parent.Clone()
	for _, opt := range options {
		opt(ctx)
	}
	return
}



func (ctx *Context) Register(hs ... Handler) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.handlers = append(ctx.handlers, hs...)
	ctx.length = len(ctx.handlers)
}

func (ctx *Context) AddBool(key string, value bool) {
	for _, h := range ctx.handlers {
		h.AddBool(key, value)
	}
}
func (ctx *Context) AddString(key string, value string) {
	for _, h := range ctx.handlers {
		h.AddString(key, value)
	}
}
func (ctx *Context) AddInt(key string, value int) {
	for _, h := range ctx.handlers {
		h.AddInt(key, value)
	}
}
func (ctx *Context) AddInt64(key string, value int64) {
	for _, h := range ctx.handlers {
		h.AddInt64(key, value)
	}
}
func (ctx *Context) AddUint(key string, value uint) {
	for _, h := range ctx.handlers {
		h.AddUint(key, value)
	}
}
func (ctx *Context) AddUint64(key string, value uint64) {
	for _, h := range ctx.handlers {
		h.AddUint64(key, value)
	}
}
func (ctx *Context) AddFloat32(key string, value float32) {
	for _, h := range ctx.handlers {
		h.AddFloat32(key, value)
	}
}
func (ctx *Context) AddFloat64(key string, value float64) {
	for _, h := range ctx.handlers {
		h.AddFloat64(key, value)
	}
}
func (ctx *Context) AddObject(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.AddObject(key, value)
	}
}
func (ctx *Context) With(fields ...KeyValue) {
	for _, h := range ctx.handlers {
		h.With(fields...)
	}
}


func (ctx *Context) Log(level Level, msg string, fields ...KeyValue) {
	if (ctx.level <= level) {
		for _, h := range ctx.handlers {
			h.Log(level, msg, fields...)
		}
	}
}
func (ctx *Context) Trace(msg string, fields ...KeyValue) {
	if (ctx.level <= TRACE) {
		for _, h := range ctx.handlers {
			h.Log(TRACE, msg, fields...)
		}
	}
}
func (ctx *Context) Debug(msg string, fields ...KeyValue) {
	if (ctx.level <= DEBUG) {
		for _, h := range ctx.handlers {
			h.Log(DEBUG, msg, fields...)
		}
	}
}
func (ctx *Context) Info(msg string, fields ...KeyValue) {
	if (ctx.level <= INFO) {
		for _, h := range ctx.handlers {
			h.Log(INFO, msg, fields...)
		}
	}
}
func (ctx *Context) Warn(msg string, fields ...KeyValue) {
	if (ctx.level <= WARNING) {
		for _, h := range ctx.handlers {
			h.Log(WARNING, msg, fields...)
		}
	}
}
func (ctx *Context) Error(msg string, fields ...KeyValue) {
	if (ctx.level <= ERROR) {
		for _, h := range ctx.handlers {
			h.Log(ERROR, msg, fields...)
		}
	}
}
func (ctx *Context) Panic(msg string, fields ...KeyValue) {
	if (ctx.level <= PANIC) {
		for _, h := range ctx.handlers {
			h.Log(PANIC, msg, fields...)
		}
	}
}
func (ctx *Context) Fatal(msg string, fields ...KeyValue) {
	if (ctx.level <= FATAL) {
		for _, h := range ctx.handlers {
			h.Log(FATAL, msg, fields...)
		}
	}
}

func (ctx *Context) Clone() (c *Context) {
	defer ctx.mu.Unlock()
	ctx.mu.Lock()

	c = &Context{
		Metadata: map[string]string{},
		mu: sync.Mutex{},
		level: ctx.level,
	}
	for _, h := range ctx.handlers {
		c.handlers = append(c.handlers, h.Clone())
	}
	for k, v := range ctx.Metadata {
		c.Metadata[k] = v
	}
	return
}



