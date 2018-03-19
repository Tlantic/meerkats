package writer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Tlantic/meerkats"
	"github.com/opentracing/opentracing-go/log"
)

var (
	partial_ts  = []byte("timestamp=\"")
	partial_lvl = []byte("\" level=\"")
	partial_msg = []byte("\" message=")
)

var pool = sync.Pool{New: func() interface{} {
	return &handler{
		Level:  meerkats.LevelAll,
		fields: make([]log.Field, 0, 8),
		fmap:   make(map[string]*log.Field),
		tl:     time.RFC3339Nano,
		w:      os.Stdout,
	}
}}

var _ meerkats.Handler = (*handler)(nil)

type handler struct {
	mu     sync.Mutex
	w      io.Writer
	tl     string
	Level  meerkats.Level
	fmap   map[string]*log.Field
	fields []log.Field
}

func (h *handler) upsertField(key string, field log.Field) {
	if f := h.fmap[key]; f != nil {
		*f = field
		return
	}
	h.fields = append(h.fields, field)
	h.fmap[key] = &h.fields[len(h.fields)-1]
	return
}

func New(options ...meerkats.HandlerOption) (h *handler) {
	h = pool.Get().(*handler)
	for _, opt := range options {
		opt.Apply(h)
	}
	return
}

func Register(options ...meerkats.HandlerOption) meerkats.LoggerOption {
	return meerkats.LoggerOptionFunc(func(l meerkats.Logger) {
		l.Register(New(options...))
	})
}

func (h *handler) Apply(l meerkats.Logger) {
	l.Register(h)
}

func (h *handler) SetLevel(level meerkats.Level) {
	h.Level = level
}
func (h *handler) GetLevel() meerkats.Level {
	return h.Level
}
func (h *handler) Close() error {
	h.Level = meerkats.LevelAll
	h.tl = time.RFC3339Nano
	h.fields = make([]log.Field, 0, 8)
	h.fmap = make(map[string]*log.Field)
	h.w = os.Stdout
	pool.Put(h)
	return nil
}
func (h *handler) Child() meerkats.Handler {
	return h.New()
}
func (h *handler) EmitBool(key string, value bool) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Bool(key, value))
}
func (h *handler) EmitString(key string, value string) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.String(key, value))
}
func (h *handler) EmitInt(key string, value int) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Int(key, value))
}
func (h *handler) EmitInt32(key string, value int32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Int32(key, value))
}
func (h *handler) EmitInt64(key string, value int64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Int64(key, value))
}
func (h *handler) EmitUint(key string, value uint) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Uint(key, value))
}
func (h *handler) EmitUint32(key string, value uint32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Uint32(key, value))
}
func (h *handler) EmitUint64(key string, value uint64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Uint64(key, value))
}
func (h *handler) EmitFloat32(key string, value float32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Float32(key, value))
}
func (h *handler) EmitFloat64(key string, value float64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Float64(key, value))
}
func (h *handler) EmitJSON(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()

	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(value)
	h.upsertField(key, meerkats.String(key, b.String()))
}
func (h *handler) EmitError(err error) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField("error", meerkats.ErrorString(err))
}
func (h *handler) EmitObject(key string, value interface{}) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.upsertField(key, meerkats.Object(key, value))
}
func (h *handler) EmitLazyLogger(value log.LazyLogger) {
	value(h)
}
func (h *handler) EmitField(fs ...log.Field) {
	for _, v := range fs {
		v.Marshal(h)
	}
}
func (h *handler) Log(t time.Time, level meerkats.Level, msg string, fields []log.Field, _ map[string]interface{}, done meerkats.DoneCallback) {
	var bs []byte
	if h.tl != "" {
		bs = t.AppendFormat(append(bs, partial_ts...), h.tl)
	}
	bs = append(append(append(append(bs, partial_lvl...), []byte(level.String())...), partial_msg...), []byte(strconv.Quote(msg))...)

	var fs []byte
	h.mu.Lock()
	for _, v := range h.fields {
		fs = append(append(fs, ' '), append(append([]byte(v.Key()), '='), []byte(strconv.Quote(fmt.Sprintf("%v", v.Value())))...)...)
	}
	h.mu.Unlock()
	for _, v := range fields {
		fs = append(append(fs, ' '), append(append([]byte(v.Key()), '='), []byte(strconv.Quote(fmt.Sprintf("%v", v.Value())))...)...)
	}
	h.w.Write(append(append(bs, fs...), '\n'))
	done()
}
