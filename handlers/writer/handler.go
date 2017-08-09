package writer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/Tlantic/meerkats"
	"github.com/opentracing/opentracing-go/log"
)

var (
	partial_lvl = []byte("level=\"")
	partial_ts  = []byte("\" timestamp=\"")
	partial_msg = []byte("\" message=\"")
)

var buffSize = 2048
var pool = sync.Pool{New: func() interface{} {
	return &writerLogger{
		Level: meerkats.LEVEL_ALL,
		tl:    time.RFC3339Nano,
		w:     os.Stdout,
		bytes: make([]byte, 0, buffSize),
	}
}}

var _ meerkats.Handler = (*writerLogger)(nil)

type writerLogger struct {
	Level meerkats.Level
	tl    string
	mu    sync.Mutex
	bytes []byte
	w     io.Writer
}

func New(options ...meerkats.HandlerOption) (h *writerLogger) {
	h = pool.Get().(*writerLogger)
	for _, opt := range options {
		opt.Apply(h)
	}
	return
}

func Register(options ...meerkats.HandlerOption) meerkats.LoggerOption {
	return meerkats.LoggerReceiver(func(l meerkats.Logger) {
		l.Register(New(options...))
	})
}

func (h *writerLogger) Apply(l meerkats.Logger) {
	l.Register(h)
}

func (h *writerLogger) SetLevel(level meerkats.Level) {
	h.Level = level
}
func (h *writerLogger) GetLevel() meerkats.Level {
	return h.Level
}
func (h *writerLogger) Dispose() {
	h.Level = meerkats.LEVEL_ALL
	h.tl = time.RFC3339Nano
	h.w = os.Stdout
	h.bytes = h.bytes[:0]
	pool.Put(h)
}
func (h *writerLogger) Clone() meerkats.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()
	clone := pool.Get().(*writerLogger)
	clone.w = h.w
	clone.bytes = h.bytes[0:]
	clone.Level = h.Level
	clone.tl = h.tl
	return clone
}

func (h *writerLogger) EmitBool(key string, value bool) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendBool(h.bytes, key, value)
}
func (h *writerLogger) AddBool(key string, value bool) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendBool(h.bytes, key, value)
}
func (h *writerLogger) EmitString(key string, value string) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, value)
}
func (h *writerLogger) AddString(key string, value string) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, value)
}
func (h *writerLogger) EmitInt(key string, value int) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, int64(value))
}
func (h *writerLogger) AddInt(key string, value int) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, int64(value))
}
func (h *writerLogger) EmitInt32(key string, value int32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, int64(value))
}
func (h *writerLogger) EmitInt64(key string, value int64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, value)
}
func (h *writerLogger) AddInt64(key string, value int64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, value)
}
func (h *writerLogger) EmitUint(key string, value uint) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, uint64(value))
}
func (h *writerLogger) AddUint(key string, value uint) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, uint64(value))
}
func (h *writerLogger) EmitUint32(key string, value uint32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, uint64(value))
}
func (h *writerLogger) EmitUint64(key string, value uint64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, value)
}
func (h *writerLogger) AddUint64(key string, value uint64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, value)
}
func (h *writerLogger) EmitFloat32(key string, value float32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat32(h.bytes, key, value)
}
func (h *writerLogger) AddFloat32(key string, value float32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat32(h.bytes, key, value)
}
func (h *writerLogger) EmitFloat64(key string, value float64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat64(h.bytes, key, value)
}
func (h *writerLogger) AddFloat64(key string, value float64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat64(h.bytes, key, value)
}
func (h *writerLogger) EmitJSON(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()

	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(value)
	h.bytes = appendString(h.bytes, key, b.String())
}
func (h *writerLogger) AddJSON(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()

	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(value)
	h.bytes = appendString(h.bytes, key, b.String())
}
func (h *writerLogger) EmitError(err error) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, "error", err.Error())
}
func (h *writerLogger) AddError(err error) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, "error", err.Error())
}
func (h *writerLogger) EmitObject(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, fmt.Sprintf("%+v", value))
}
func (h *writerLogger) EmitLazyLogger(value log.LazyLogger) {
	value(h)
}
func (h *writerLogger) Add(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, fmt.Sprintf("%+v", value))
}

func (h *writerLogger) EmitField(fs ...log.Field) {
	for _, v := range fs {
		v.Marshal(h)
	}
}
func (h *writerLogger) With(fs ...meerkats.Field) {
	for _, v := range fs {
		switch v.Type {
		case meerkats.TypeString:
			h.AddString(v.Key, v.ValueString)
		case meerkats.TypeBool:
			h.AddBool(v.Key, v.ValueBool)
		case meerkats.TypeInt64:
			h.AddInt64(v.Key, v.ValueInt64)
		case meerkats.TypeUint64:
			h.AddUint64(v.Key, v.ValueUint64)
		case meerkats.TypeFloat32:
			h.AddFloat32(v.Key, v.ValueFloat32)
		case meerkats.TypeFloat64:
			h.AddFloat64(v.Key, v.ValueFloat64)
		case meerkats.TypeError:
			h.AddError(v.ValueInterface.(error))
		case meerkats.TypeJSON:
			h.AddJSON(v.Key, v.ValueInterface)
		default:
			h.Add(v.Key, v.ValueInterface)
		}
	}
}

func (h *writerLogger) Log(t time.Time, level meerkats.Level, msg string, fields []log.Field, _ map[string]interface{}, done meerkats.DoneCallback) {
	clone := pool.Get().(*writerLogger)
	clone.bytes = append(append(append(clone.bytes, partial_lvl...), level.String()...))
	if h.tl != "" {
		clone.bytes = append(clone.bytes, partial_ts...)
		clone.bytes = t.AppendFormat(clone.bytes, h.tl)
	}
	clone.bytes = append(append(append(append(clone.bytes, partial_msg...), msg...), '"'), h.bytes[0:]...)
	clone.EmitField(fields...)
	h.w.Write(append(clone.bytes, '\n'))
	clone.Dispose()
	done()
}
