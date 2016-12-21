package meerkats

import (
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &context{
			handlers: make([]Handler, 0, 1),
		}
	},
}

var _ Logger = (*context)(nil)
type context struct {
	metadata map[string]string
	Level    Level
	handlers []Handler
	mu       sync.Mutex
	wg       sync.WaitGroup
}

func New(options...LoggerOption) Logger {
	ctx := pool.Get().(*context)
	ctx.metadata = make(map[string]string)
	ctx.Level = TRACE
	for _, opt := range options {
		opt.Apply(ctx)
	}
	return ctx
}
func From( parent Logger, options...LoggerOption) ( ctx  Logger) {
	ctx = parent.Clone()
	for _, opt := range options {
		opt.Apply(ctx)
	}
	return
}


func (ctx *context) SetLevel(lvl Level) {
	ctx.Level = lvl
}

func (ctx *context) Register(hs ... Handler) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.handlers = append(ctx.handlers, hs...)
}

func (ctx *context) SetMeta(key string, value string) {
	ctx.metadata[key] = value
}
func (ctx *context) GetMeta(key string) string {
	return ctx.metadata[key]
}


func (ctx *context) AddBool(key string, value bool) {
	for _, h := range ctx.handlers {
		h.AddBool(key, value)
	}
}
func (ctx *context) AddString(key string, value string) {
	for _, h := range ctx.handlers {
		h.AddString(key, value)
	}
}
func (ctx *context) AddInt(key string, value int) {
	for _, h := range ctx.handlers {
		h.AddInt(key, value)
	}
}
func (ctx *context) AddInt64(key string, value int64) {
	for _, h := range ctx.handlers {
		h.AddInt64(key, value)
	}
}
func (ctx *context) AddUint(key string, value uint) {
	for _, h := range ctx.handlers {
		h.AddUint(key, value)
	}
}
func (ctx *context) AddUint64(key string, value uint64) {
	for _, h := range ctx.handlers {
		h.AddUint64(key, value)
	}
}
func (ctx *context) AddFloat32(key string, value float32) {
	for _, h := range ctx.handlers {
		h.AddFloat32(key, value)
	}
}
func (ctx *context) AddFloat64(key string, value float64) {
	for _, h := range ctx.handlers {
		h.AddFloat64(key, value)
	}
}
func (ctx *context) Add(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.Add(key, value)
	}
}
func (ctx *context) With(fields ...Field) {
	for _, h := range ctx.handlers {
		h.With(fields...)
	}
}


func (ctx *context) Log(level Level, msg string, fields ...Field) {
	if (ctx.Level <= level) {
		now := time.Now()
		for _, h := range ctx.handlers {
			h.Log(now, level, msg, fields, ctx.metadata)
		}
	}
}
func (ctx *context) Trace(msg string, fields ...Field) {
	if (ctx.Level <= TRACE) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&TRACE != 0 ) {
				h.Log(now, TRACE, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Debug(msg string, fields ...Field) {
	if (ctx.Level <= DEBUG) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&DEBUG != 0 ) {
				h.Log(now, DEBUG, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Info(msg string, fields ...Field) {
	if ctx.Level <= INFO {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&INFO != 0 ) {
				h.Log(now, INFO, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Warn(msg string, fields ...Field) {
	if (ctx.Level <= WARNING) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&WARNING != 0 ) {
				h.Log(now, WARNING, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Error(msg string, fields ...Field) {
	if (ctx.Level <= ERROR) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&ERROR != 0 ) {
				h.Log(now, ERROR, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Panic(msg string, fields ...Field) {
	if (ctx.Level <= PANIC) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&PANIC != 0 ) {
				h.Log(now, PANIC, msg, fields, ctx.metadata)
			}
		}
	}
}
func (ctx *context) Fatal(msg string, fields ...Field) {
	if (ctx.Level <= FATAL) {
		now := time.Now()
		for _, h := range ctx.handlers {
			if ( h.GetLevel()&FATAL != 0 ) {
				h.Log(now, FATAL, msg, fields, ctx.metadata)
			}
		}
	}
}


func (ctx *context) Clone() Logger {
	defer ctx.mu.Unlock()
	ctx.mu.Lock()
	c := pool.Get().(*context)
	c.handlers = ctx.handlers
	c.metadata = map[string]string{}
	c.Level = ctx.Level
	for _, h := range ctx.handlers {
		c.handlers = append(c.handlers, h.Clone())
	}
	for k, v := range ctx.metadata {
		c.metadata[k] = v
	}
	return c
}
func (ctx *context) Dispose() () {
	ctx.handlers = ctx.handlers[:0]
	pool.Put(ctx)
}



