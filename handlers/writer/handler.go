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

var buffSize = 2048
var pool = sync.Pool{New: func() interface{} {
	return &writerLogger{
		Level: meerkats.LEVEL_ALL,
		timelayout: time.RFC3339Nano,
		w: os.Stdout,
		bytes: make([]byte, 0, buffSize),
	}
}}

var _ meerkats.Handler = (*writerLogger)(nil)
type writerLogger struct {
	Level      meerkats.Level
	timelayout string
	mu         sync.Mutex
	bytes      []byte
	w          io.Writer
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
	h.timelayout = time.RFC3339Nano
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
	clone.timelayout = h.timelayout
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
func (h *writerLogger) Add(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.bytes = appendString(h.bytes, key, fmt.Sprintf("%+v", value))
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
		default:
			h.Add(v.Key, v.ValueInterface)
		}
	}
}


func (h *writerLogger) Log(t time.Time, level meerkats.Level, msg string, fields []meerkats.Field, _ map[string]string) {
	clone := pool.Get().(*writerLogger)
	clone.bytes = append(append(append(clone.bytes, partial_lvl...), level.String()...))
	if (h.timelayout != "") {
		clone.bytes = append(clone.bytes, partial_ts...)
		clone.bytes = t.AppendFormat(clone.bytes, h.timelayout)
	}
	clone.bytes = append(append(append(append(clone.bytes, partial_msg...), msg...), '"'), h.bytes[0:]...)
	clone.With(fields...)
	h.w.Write(append(clone.bytes, '\n'))
	clone.Dispose()
}