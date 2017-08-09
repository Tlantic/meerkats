package meerkats

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"os"
	"regexp"
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

type ctxKey uint8

const (
	uniqueKey ctxKey = iota
)

type context struct {
	handlerMu   sync.Mutex
	spanMu      sync.Mutex
	wg          sync.WaitGroup
	opName      string
	span        opentracing.Span
	Level       Level
	writerLevel Level
	metadata    map[string]interface{}
	handlers    []Handler
}

func New(options ...LoggerOption) Logger {
	ctx := pool.Get().(*context)
	ctx.metadata = make(map[string]interface{})
	ctx.Level = TRACE
	ctx.writerLevel = TRACE

	for _, opt := range append(append(([]LoggerOption)(nil), newSpanHandler()), options...) {
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
	ctx.span.SetOperationName(name)
	ctx.opName = name
}
func (ctx *context) Span() (span opentracing.Span) {
	ctx.spanMu.Lock()
	defer ctx.spanMu.Unlock()
	if ctx.span == nil {
		ctx.span = opentracing.StartSpan(ctx.OperationName())
		for k, v := range ctx.metadata {
			ctx.span.SetTag(k, v)
		}
	}
	return ctx.span
}
func (ctx *context) WithSpan(span opentracing.Span) {
	if span != nil {
		span = span.Tracer().StartSpan(ctx.OperationName(),
			opentracing.SpanReference{
				Type:              opentracing.FollowsFromRef,
				ReferencedContext: ctx.Span().Context(),
			}, opentracing.SpanReference{
				Type:              opentracing.ChildOfRef,
				ReferencedContext: span.Context(),
			})

		ctx.spanMu.Lock()
		for k, v := range ctx.metadata {
			span.SetTag(k, v)
		}
		ctx.span = span
		ctx.spanMu.Unlock()
	}
}
func (ctx *context) SetLevel(lvl Level) {
	ctx.Level = lvl
}

func (ctx *context) Register(hs ...Handler) {
	ctx.handlerMu.Lock()
	defer ctx.handlerMu.Unlock()
	ctx.handlers = append(ctx.handlers, hs...)
}

// Deprecate: Use SetTag
func (ctx *context) SetMeta(key string, value string) {
	ctx.SetTag(key, value)
}
func (ctx *context) SetTag(key string, value interface{}) {
	ctx.Span().SetTag(key, value)
	ctx.metadata[key] = value
}

// Deprecate: Use GetTag
func (ctx *context) GetMeta(key string) string {
	s, _ := ctx.metadata[key].(string)
	return s
}
func (ctx *context) GetTag(key string) interface{} {
	return ctx.metadata[key]
}

func (ctx *context) EmitBool(key string, value bool) {
	for _, h := range ctx.handlers {
		h.EmitBool(key, value)
	}
}
func (ctx *context) AddBool(key string, value bool) {
	for _, h := range ctx.handlers {
		h.EmitBool(key, value)
	}
}

func (ctx *context) EmitString(key string, value string) {
	for _, h := range ctx.handlers {
		h.EmitString(key, value)
	}
}
func (ctx *context) AddString(key string, value string) {
	for _, h := range ctx.handlers {
		h.EmitString(key, value)
	}
}

func (ctx *context) EmitInt(key string, value int) {
	for _, h := range ctx.handlers {
		h.EmitInt(key, value)
	}
}
func (ctx *context) AddInt(key string, value int) {
	for _, h := range ctx.handlers {
		h.EmitInt(key, value)
	}
}

func (ctx *context) EmitInt32(key string, value int32) {
	for _, h := range ctx.handlers {
		h.EmitInt32(key, value)
	}
}

func (ctx *context) EmitInt64(key string, value int64) {
	for _, h := range ctx.handlers {
		h.EmitInt64(key, value)
	}
}
func (ctx *context) AddInt64(key string, value int64) {
	for _, h := range ctx.handlers {
		h.EmitInt64(key, value)
	}
}

func (ctx *context) EmitUint(key string, value uint) {
	for _, h := range ctx.handlers {
		h.EmitUint(key, value)
	}
}
func (ctx *context) AddUint(key string, value uint) {
	for _, h := range ctx.handlers {
		h.EmitUint(key, value)
	}
}
func (ctx *context) EmitUint32(key string, value uint32) {
	for _, h := range ctx.handlers {
		h.EmitUint32(key, value)
	}
}
func (ctx *context) EmitUint64(key string, value uint64) {
	for _, h := range ctx.handlers {
		h.EmitUint64(key, value)
	}
}
func (ctx *context) AddUint64(key string, value uint64) {
	for _, h := range ctx.handlers {
		h.EmitUint64(key, value)
	}
}

func (ctx *context) EmitFloat32(key string, value float32) {
	for _, h := range ctx.handlers {
		h.EmitFloat32(key, value)
	}
}
func (ctx *context) AddFloat32(key string, value float32) {
	for _, h := range ctx.handlers {
		h.EmitFloat32(key, value)
	}
}

func (ctx *context) EmitFloat64(key string, value float64) {
	for _, h := range ctx.handlers {
		h.EmitFloat64(key, value)
	}
}
func (ctx *context) AddFloat64(key string, value float64) {
	for _, h := range ctx.handlers {
		h.EmitFloat64(key, value)
	}
}

func (ctx *context) EmitJSON(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.EmitJSON(key, value)
	}
}
func (ctx *context) AddJSON(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.EmitJSON(key, value)
	}
}

func (ctx *context) EmitError(err error) {
	for _, h := range ctx.handlers {
		h.EmitError(err)
	}
}
func (ctx *context) AddError(err error) {
	for _, h := range ctx.handlers {
		h.EmitError(err)
	}
}

func (ctx *context) EmitObject(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.EmitObject(key, value)
	}
}
func (ctx *context) Add(key string, value interface{}) {
	for _, h := range ctx.handlers {
		h.EmitObject(key, value)
	}
}

func (ctx *context) With(fields ...Field) {
	for _, h := range ctx.handlers {
		h.With(fields...)
	}
}
func (ctx *context) EmitField(fields ...log.Field) {
	for _, h := range ctx.handlers {
		h.EmitField(fields...)
	}
}

func (ctx *context) EmitLazyLogger(value log.LazyLogger) {
	for _, h := range ctx.handlers {
		h.EmitLazyLogger(value)
	}
}

func (ctx *context) Log(level Level, msg string, fields ...log.Field) {
	if ctx.Level <= level {
		now := time.Now()
		for _, h := range ctx.handlers {
			ctx.wg.Add(1)
			h.Log(now, level, msg, fields, ctx.metadata, ctx.wg.Done)
		}
	}
}
func (ctx *context) Trace(msg string, fields ...log.Field) {
	if ctx.Level <= TRACE {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&TRACE != 0 {
				ctx.wg.Add(1)
				h.Log(now, TRACE, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
}
func (ctx *context) Debug(msg string, fields ...log.Field) {
	if ctx.Level <= DEBUG {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&DEBUG != 0 {
				ctx.wg.Add(1)
				h.Log(now, DEBUG, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
}
func (ctx *context) Info(msg string, fields ...log.Field) {
	if ctx.Level <= INFO {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&INFO != 0 {
				ctx.wg.Add(1)
				h.Log(now, INFO, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
}
func (ctx *context) Warn(msg string, fields ...log.Field) {
	if ctx.Level <= WARNING {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&WARNING != 0 {
				ctx.wg.Add(1)
				h.Log(now, WARNING, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
}
func (ctx *context) Error(msg string, fields ...log.Field) {
	if ctx.Level <= ERROR {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&ERROR != 0 {
				ctx.wg.Add(1)
				h.Log(now, ERROR, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
}
func (ctx *context) Panic(msg string, fields ...log.Field) {
	if ctx.Level <= PANIC {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&PANIC != 0 {
				ctx.wg.Add(1)
				h.Log(now, PANIC, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
	ctx.Dispose()
	panic(msg)
}
func (ctx *context) Fatal(msg string, fields ...log.Field) {
	if ctx.Level <= FATAL {
		now := time.Now()
		for _, h := range ctx.handlers {
			if h.GetLevel()&FATAL != 0 {
				ctx.wg.Add(1)
				h.Log(now, FATAL, msg, fields, ctx.metadata, ctx.wg.Done)
			}
		}
	}
	ctx.Dispose()
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
	c.metadata = map[string]interface{}{}

	ctx.handlerMu.Lock()
	defer ctx.handlerMu.Unlock()

	c.Level = ctx.Level
	if ctx.span != nil {
		c.span = ctx.span.Tracer().StartSpan(ctx.OperationName(), opentracing.SpanReference{
			ReferencedContext: ctx.span.Context(),
			Type:              opentracing.ChildOfRef,
		})
	}

	for _, h := range ctx.handlers {
		h.Clone().Apply(c)
	}
	for k, v := range ctx.metadata {
		c.SetTag(k, v)
	}
	return c
}
func (ctx *context) Dispose() {
	ctx.wg.Wait()
	ctx.span.Finish()
	ctx.handlers = ctx.handlers[:0]
	pool.Put(ctx)
}
