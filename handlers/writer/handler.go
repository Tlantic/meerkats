package writer

import (
	"io"
	"os"
	"fmt"
	"time"
	"sync"

	"github.com/Tlantic/meerkats"
)

var (
	partial_lvl = []byte("level=\"")
	partial_ts = []byte("\" timestamp=\"")
	partial_msg = []byte("\" message=\"")
)

var buffSize = 1024
var pool = sync.Pool{New: func() interface{} {
	return &writerLogger{
		bytes: make([]byte, 0, buffSize),
	}
}}

type writerLogger struct {
	Level      meerkats.Level
	TimeLayout string
	mu         sync.Mutex
	bytes      []byte
	w          io.Writer
}

func New(options ...meerkats.HandlerOption) (h *writerLogger) {
	h = pool.Get().(*writerLogger)
	h.bytes = h.bytes[:0]
	h.w = os.Stdout
	h.TimeLayout = time.RFC3339Nano
	h.Level = meerkats.LEVEL_ALL
	for _, opt := range options {
		opt(h)
	}
	return
}

func Register(options ...meerkats.HandlerOption) meerkats.ContextOption {
	return meerkats.ContextOption(func(ctx *meerkats.Context) {
		ctx.Register(New(options...))
	})
}



func (h *writerLogger) SetTimeLayout(layout string) {
	h.TimeLayout = layout
}
func (h *writerLogger) GetTimeLayout() string {
	return h.TimeLayout
}

func (h *writerLogger) SetLevel(level meerkats.Level) {
	h.Level = level
}
func (h *writerLogger) GetLevel() meerkats.Level {
	return h.Level
}
func (h *writerLogger) Dispose() {
	pool.Put(h)
}
func (h *writerLogger) Clone() meerkats.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	clone := New()
	clone.w = h.w
	clone.Level = h.Level
	clone.TimeLayout = h.TimeLayout
	clone.bytes = h.bytes[0:]
	return clone
}


func (h *writerLogger) AddBool(key string, value bool) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendBool(h.bytes, key, value)
}

func (h *writerLogger) AddString(key string, value string){
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, value)
}
func (h *writerLogger) AddInt(key string, value int) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, int64(value))
}
func (h *writerLogger) AddInt64(key string, value int64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendInt64(h.bytes, key, value)
}
func (h *writerLogger) AddUint(key string, value uint) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, uint64(value))
}
func (h *writerLogger) AddUint64(key string, value uint64)  {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendUint64(h.bytes, key, value)
}
func (h *writerLogger) AddFloat32(key string, value float32)  {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat32(h.bytes, key, value)
}
func (h *writerLogger) AddFloat64(key string, value float64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendFloat64(h.bytes, key, value)
}
func (h *writerLogger) AddObject(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, fmt.Sprintf("%+v", value))
}

func (h *writerLogger) With(fs ...meerkats.KeyValue) {
	for _, v := range fs {
		switch v.GetType() {
		case meerkats.TypeString:
			h.AddString(v.Key, v.GetString())
		case meerkats.TypeBool:
			h.AddBool(v.Key, v.GetBool())
		case meerkats.TypeInt64:
			h.AddInt64(v.Key, v.GetInt64())
		case meerkats.TypeUint64:
			h.AddUint64(v.Key, v.GetUint64())
		case meerkats.TypeFloat32:
			h.AddFloat32(v.Key, v.GetFloat32())
		case meerkats.TypeFloat64:
			h.AddFloat64(v.Key, v.GetFloat64())
		default:
			h.AddObject(v.Key, v.GetInterface())
		}
	}
}


func (h *writerLogger) Log(level meerkats.Level, msg string, fields []meerkats.KeyValue) {
	if ( h.Level&level != 0 ) {
		log(h.w, level, time.Now(), h.TimeLayout, msg, h.bytes[0:], fields)
	}
}


func log(w io.Writer, level meerkats.Level, t time.Time, format string, msg string, pre []byte, fields []meerkats.KeyValue) {
		clone := pool.Get().(*writerLogger)
		clone.bytes = append(append(append(clone.bytes[:0], partial_lvl...), level.String()...))
		if (format != "") {
			clone.bytes = append(clone.bytes, partial_ts...)
			clone.bytes = t.AppendFormat(clone.bytes, format)

		}
		clone.bytes = append(append(append(append(clone.bytes, partial_msg...), msg...), '"'), pre...)
		clone.With(fields...)
		w.Write(append(clone.bytes, '\n'))
		clone.Dispose()
}