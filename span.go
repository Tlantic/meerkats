package meerkats

import (
	"bytes"
	"encoding/json"
	"github.com/opentracing/opentracing-go/log"
	"sync"
	"time"
)

type spanHandler struct {
	mu     sync.Mutex
	Level  Level
	logger Logger
	fields []log.Field
}

func newSpanHandler(options ...HandlerOption) Handler {
	s := &spanHandler{
		Level: TRACE,
	}
	for _, v := range options {
		v.Apply(s)
	}
	return s
}

func (h *spanHandler) Apply(l Logger) {
	h.logger = l
	l.Register(h)
}

func (h *spanHandler) SetLevel(level Level) {
	h.Level = level
}
func (h *spanHandler) GetLevel() Level {
	return h.Level
}
func (h *spanHandler) Dispose() {
	pool.Put(h)
}
func (h *spanHandler) Clone() Handler {
	return &spanHandler{
		sync.Mutex{},
		h.Level,
		h.logger,
		h.fields[:],
	}
}

func (h *spanHandler) EmitBool(key string, value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Bool(key, value))
}
func (h *spanHandler) AddBool(key string, value bool) {
	h.EmitBool(key, value)
}
func (h *spanHandler) EmitString(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, String(key, value))
}
func (h *spanHandler) AddString(key string, value string) {
	h.EmitString(key, value)
}
func (h *spanHandler) EmitInt(key string, value int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Int(key, value))
}
func (h *spanHandler) AddInt(key string, value int) {
	h.EmitInt(key, value)
}
func (h *spanHandler) EmitInt32(key string, value int32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Int32(key, value))
}
func (h *spanHandler) EmitInt64(key string, value int64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Int64(key, value))
}
func (h *spanHandler) AddInt64(key string, value int64) {
	h.EmitInt64(key, value)
}
func (h *spanHandler) EmitUint(key string, value uint) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Uint(key, value))
}
func (h *spanHandler) AddUint(key string, value uint) {
	h.EmitUint(key, value)
}

func (h *spanHandler) EmitUint32(key string, value uint32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Uint32(key, value))
}

func (h *spanHandler) EmitUint64(key string, value uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Uint64(key, value))
}
func (h *spanHandler) AddUint64(key string, value uint64) {
	h.EmitUint64(key, value)
}

func (h *spanHandler) EmitFloat32(key string, value float32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Float32(key, value))
}
func (h *spanHandler) AddFloat32(key string, value float32) {
	h.EmitFloat32(key, value)
}

func (h *spanHandler) EmitFloat64(key string, value float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Float64(key, value))
}
func (h *spanHandler) AddFloat64(key string, value float64) {
	h.EmitFloat64(key, value)
}

func (h *spanHandler) EmitJSON(key string, value interface{}) {
	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(value)
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, String(key, b.String()))
}
func (h *spanHandler) AddJSON(key string, value interface{}) {
	h.EmitJSON(key, value)
}

func (h *spanHandler) AddError(err error) {
	h.EmitError(err)
}
func (h *spanHandler) EmitError(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, String("error", err.Error()))
}
func (h *spanHandler) EmitObject(key string, value interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, Object(key, value))
}
func (h *spanHandler) EmitLazyLogger(value log.LazyLogger) {
	value(h)
}
func (h *spanHandler) Add(key string, value interface{}) {
	h.EmitObject(key, value)
}

func (h *spanHandler) With(fs ...Field) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, v := range fs {
		h.fields = append(h.fields, Object(v.Key, v.Get()))
	}
}
func (h *spanHandler) EmitField(fs ...log.Field) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields = append(h.fields, fs...)
}

func (h *spanHandler) Log(t time.Time, level Level, msg string, fields []log.Field, _ map[string]interface{}, done DoneCallback) {
	defer done()
	if level == 0 || h.Level >= level {
		h.logger.Span().LogFields(append(append(append(([]log.Field)(nil), String("level", level.String()), String("message", msg)), h.fields...), fields...)...)
	}
}
