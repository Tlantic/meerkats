package meerkats

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"os"
	"regexp"
	"sync"
	"time"
)

type ctxKey uint8

const (
	uniqueKey ctxKey = iota
)

type metadata struct {
	sync.RWMutex
	kv map[string]interface{}
}

func (m *metadata) clear() {
	m.kv = map[string]interface{}{}
}
func (m *metadata) get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()
	return m.kv[key]
}
func (m *metadata) set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.kv[key] = value
}
func (m *metadata) forEach(fn func(key string, value interface{})) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.kv {
		fn(k, v)
	}
}

type handlerCollection struct {
	sync.RWMutex
	col []Handler
}

func (c *handlerCollection) clear() {
	c.Lock()
	for _, v := range c.col {
		v.Dispose()
	}
	c.col = nil
	c.Unlock()
}
func (c *handlerCollection) forEach(fn func(idx int, value Handler)) {
	c.RLock()
	defer c.RUnlock()
	for i, v := range c.col {
		fn(i, v)
	}
}
func (c *handlerCollection) add(values ...Handler) {
	c.Lock()
	defer c.Unlock()
	c.col = append(c.col, values...)
}

type span struct {
	sync.RWMutex
	opentracing.Span
}

func (s *span) clear() {
	s.Lock()
	defer s.Unlock()
	if s.Span != nil {
		s.Span.Finish()
	}
	s.Span = nil
}
func (s *span) setTag(key string, value interface{}) {
	s.Span.SetTag(key, value)
}

var pool = sync.Pool{
	New: func() interface{} {
		return &context{
			Level:    TRACE,
			metadata: metadata{kv: map[string]interface{}{}},
			handlers: handlerCollection{col: nil},
		}
	},
}

var _ Logger = (*context)(nil)

type context struct {
	wg       sync.WaitGroup
	opName   string
	span     span
	Level    Level
	metadata metadata
	handlers handlerCollection
}

func New(options ...LoggerOption) Logger {
	ctx := pool.Get().(*context)
	for _, opt := range append(options, newSpanHandler()) {
		opt.Apply(ctx)
	}
	return ctx
}
func From(parent Logger, options ...LoggerOption) (ctx Logger) {
	ctx = parent.Clone()
	for _, opt := range options {
		opt.Apply(ctx)
	}
	return
}

func (ctx *context) OperationName() string {
	return ctx.opName
}
func (ctx *context) SetOperationName(name string) {
	ctx.opName = name
	ctx.Span().SetOperationName(name)
}
func (ctx *context) Span() (span opentracing.Span) {
	ctx.span.Lock()
	defer ctx.span.Unlock()

	if ctx.span.Span == nil {
		ctx.span.Span = opentracing.StartSpan(ctx.OperationName())
		ctx.metadata.forEach(ctx.span.setTag)
	}
	return ctx.span.Span
}
func (ctx *context) WithSpan(span opentracing.Span) {
	if span != nil {
		ctx.span.Lock()
		defer ctx.span.Unlock()
		if s := ctx.span.Span; s != nil {
			ctx.span.Span = s.Tracer().StartSpan(ctx.OperationName(),
				opentracing.SpanReference{
					Type:              opentracing.FollowsFromRef,
					ReferencedContext: s.Context(),
				}, opentracing.SpanReference{
					Type:              opentracing.ChildOfRef,
					ReferencedContext: span.Context(),
				})
		}
		ctx.metadata.forEach(ctx.span.setTag)
	}
}
func (ctx *context) SetLevel(lvl Level) {
	ctx.Level = lvl
}

func (ctx *context) Register(hs ...Handler) {
	ctx.handlers.add(hs...)
}

// Deprecate: Use Tag
func (ctx *context) SetMeta(key string, value string) {
	ctx.SetTag(key, value)
}
func (ctx *context) SetTag(key string, value interface{}) {
	ctx.Span().SetTag(key, value)
	ctx.metadata.set(key, value)
}

// Deprecate: Use GetTag
func (ctx *context) GetMeta(key string) string {
	s, _ := ctx.GetTag(key).(string)
	return s
}
func (ctx *context) GetTag(key string) interface{} {
	return ctx.metadata.get(key)
}

func (ctx *context) EmitBool(key string, value bool) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitBool(key, value)
	})
}
func (ctx *context) AddBool(key string, value bool) {
	ctx.EmitBool(key, value)
}

func (ctx *context) EmitString(key string, value string) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitString(key, value)
	})
}
func (ctx *context) AddString(key string, value string) {
	ctx.EmitString(key, value)
}

func (ctx *context) EmitInt(key string, value int) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitInt(key, value)
	})
}
func (ctx *context) AddInt(key string, value int) {
	ctx.EmitInt(key, value)
}

func (ctx *context) EmitInt32(key string, value int32) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitInt32(key, value)
	})
}

func (ctx *context) EmitInt64(key string, value int64) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitInt64(key, value)
	})
}
func (ctx *context) AddInt64(key string, value int64) {
	ctx.EmitInt64(key, value)
}
func (ctx *context) EmitUint(key string, value uint) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitUint(key, value)
	})
}
func (ctx *context) AddUint(key string, value uint) {
	ctx.EmitUint(key, value)
}
func (ctx *context) EmitUint32(key string, value uint32) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitUint32(key, value)
	})
}
func (ctx *context) EmitUint64(key string, value uint64) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitUint64(key, value)
	})
}
func (ctx *context) AddUint64(key string, value uint64) {
	ctx.EmitUint64(key, value)
}

func (ctx *context) EmitFloat32(key string, value float32) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitFloat32(key, value)
	})
}
func (ctx *context) AddFloat32(key string, value float32) {
	ctx.EmitFloat32(key, value)
}

func (ctx *context) EmitFloat64(key string, value float64) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitFloat64(key, value)
	})
}
func (ctx *context) AddFloat64(key string, value float64) {
	ctx.EmitFloat64(key, value)
}

func (ctx *context) EmitJSON(key string, value interface{}) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitJSON(key, value)
	})
}
func (ctx *context) AddJSON(key string, value interface{}) {
	ctx.EmitJSON(key, value)
}

func (ctx *context) EmitError(err error) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitError(err)
	})
}
func (ctx *context) AddError(err error) {
	ctx.EmitError(err)
}

func (ctx *context) EmitObject(key string, value interface{}) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitObject(key, value)
	})
}
func (ctx *context) Add(key string, value interface{}) {
	ctx.EmitObject(key, value)
}

func (ctx *context) With(fields ...Field) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.With(fields...)
	})
}
func (ctx *context) EmitField(fields ...log.Field) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitField(fields...)
	})
}

func (ctx *context) EmitLazyLogger(value log.LazyLogger) {
	ctx.handlers.forEach(func(_ int, h Handler) {
		h.EmitLazyLogger(value)
	})
}

func (ctx *context) Log(level Level, msg string, fields ...log.Field) {
	if ctx.Level <= level {
		now := time.Now()
		ctx.handlers.forEach(func(_ int, h Handler) {
			if h.GetLevel()&level != 0 {
				ctx.wg.Add(1)
				h.Log(now, level, msg, fields, ctx.metadata.kv, ctx.wg.Done)
			}
		})
	}
}
func (ctx *context) Trace(msg string, fields ...log.Field) {
	ctx.Log(TRACE, msg, fields...)
}
func (ctx *context) Debug(msg string, fields ...log.Field) {
	ctx.Log(DEBUG, msg, fields...)
}
func (ctx *context) Info(msg string, fields ...log.Field) {
	ctx.Log(INFO, msg, fields...)
}
func (ctx *context) Warn(msg string, fields ...log.Field) {
	ctx.Log(WARNING, msg, fields...)
}
func (ctx *context) Error(msg string, fields ...log.Field) {
	ctx.Log(ERROR, msg, fields...)
}
func (ctx *context) Panic(msg string, fields ...log.Field) {
	defer ctx.Dispose()
	ctx.Log(PANIC, msg, fields...)
	panic(msg)
}
func (ctx *context) Fatal(msg string, fields ...log.Field) {
	defer ctx.Dispose()
	ctx.Log(FATAL, msg, fields...)
	os.Exit(1)
}

var _reNewline = regexp.MustCompile(`\r?\n`)

func (ctx *context) Write(p []byte) (n int, err error) {
	n = len(p)
	ctx.Log(ctx.Level, _reNewline.ReplaceAllString(string(p), ""))
	return
}
func (ctx *context) Clone() Logger {
	c := pool.Get().(*context)
	c.Level = ctx.Level

	ctx.span.Lock()
	defer ctx.span.Unlock()
	if s := ctx.span.Span; s != nil {
		c.span.Span = s.Tracer().StartSpan(ctx.OperationName(), opentracing.SpanReference{
			ReferencedContext: s.Context(),
			Type:              opentracing.ChildOfRef,
		})
	}

	ctx.handlers.forEach(func(_ int, h Handler) { h.Clone().Apply(c) })
	ctx.metadata.forEach(c.SetTag)

	return c
}
func (ctx *context) Dispose() {
	ctx.Level = TRACE

	ctx.wg.Wait()
	ctx.span.clear()
	ctx.handlers.clear()
	ctx.metadata.clear()
	pool.Put(ctx)
}
