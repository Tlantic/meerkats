package meerkats

import (
	"bytes"
	"encoding/json"
	"github.com/opentracing/opentracing-go/log"
	"sync"
	"time"
)

var spool = sync.Pool{
	New: func() interface{} {
		return &spanHandler{
			Level:  LevelAll,
			fields: map[string]log.Field{},
		}
	},
}

type spanHandler struct {
	mu     sync.Mutex
	Level  Level
	logger Logger
	fields map[string]log.Field
}

func newSpanHandler(options ...HandlerOption) Handler {
	s := spool.Get().(*spanHandler)
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
func (h *spanHandler) Close() error {
	h.logger = nil
	h.fields = map[string]log.Field{}
	h.Level = LevelAll
	spool.Put(h)
	return nil
}
func (h *spanHandler) Child() Handler {
	h2 := spool.Get().(*spanHandler)
	h2.Level = h.Level
	h2.logger = h.logger
	for k, v := range h.fields {
		h2.fields[k] = v
	}
	return h2
}

func (h *spanHandler) EmitBool(key string, value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Bool(key, value)
}
func (h *spanHandler) EmitString(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = String(key, value)
}
func (h *spanHandler) EmitInt(key string, value int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Int(key, value)
}
func (h *spanHandler) EmitInt32(key string, value int32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Int32(key, value)
}
func (h *spanHandler) EmitInt64(key string, value int64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Int64(key, value)
}
func (h *spanHandler) EmitUint(key string, value uint) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Uint(key, value)
}
func (h *spanHandler) EmitUint32(key string, value uint32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Uint32(key, value)
}
func (h *spanHandler) EmitUint64(key string, value uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Uint64(key, value)
}
func (h *spanHandler) EmitFloat32(key string, value float32) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Float32(key, value)
}
func (h *spanHandler) EmitFloat64(key string, value float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Float64(key, value)
}
func (h *spanHandler) EmitJSON(key string, value interface{}) {
	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(value)
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = String(key, b.String())
}
func (h *spanHandler) EmitError(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields["error"] = String("error", err.Error())
}
func (h *spanHandler) EmitObject(key string, value interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[key] = Object(key, value)
}
func (h *spanHandler) EmitLazyLogger(value log.LazyLogger) {
	value(h)
}
func (h *spanHandler) EmitField(fs ...log.Field) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, f := range fs {
		h.fields[f.Key()] = f
	}
}
func (h *spanHandler) Log(t time.Time, level Level, msg string, fields []log.Field, _ map[string]interface{}, done DoneCallback) {
	defer done()
	if level == 0 || h.Level >= level {
		var fs []log.Field
		h.mu.Lock()
		for _, f := range h.fields {
			fs = append(fs, f)
		}
		h.mu.Unlock()

		h.logger.Span().LogFields(append(append(([]log.Field)(nil), String("level", level.String()), String("message", msg)), append(fs, fields...)...)...)
	}
}
