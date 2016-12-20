package stdlogger

import (
	"fmt"
	"log"
	"time"
	"sync"
	"bytes"
	"strconv"

	"github.com/Tlantic/meerkats"
)

var buffSize = 256
var pool = sync.Pool{New: func() interface{} {
	return &stdLogger{
		bytes: make([]byte, 0, buffSize),
	}
}}

//noinspection GoUnusedConst
const (
	COLOR_NONE = "0"
	COLOR_RED = "31"
	COLOR_GREEN = "32"
	COLOR_YELLOW = "33"
	COLOR_BLUE = "34"
	COLOR_GRAY = "37"
)

var COLORS = [...]string{
	meerkats.TRACE: COLOR_NONE,
	meerkats.DEBUG: COLOR_GRAY,
	meerkats.INFO:  COLOR_BLUE,
	meerkats.WARNING:  COLOR_YELLOW,
	meerkats.ERROR: COLOR_RED,
	meerkats.FATAL: COLOR_RED,
	meerkats.PANIC: COLOR_RED,
}

type stdLogger struct {
	Level      	meerkats.Level
	TimeLayout 	string
	mu         	sync.Mutex
	bytes      	[]byte
}

func New(options ...meerkats.HandlerOption) (h *stdLogger) {
	h = pool.Get().(*stdLogger)
	h.TimeLayout = time.RFC3339Nano
	h.Level = meerkats.LEVEL_ALL
	for _, opt := range options {
		opt( h )
	}
	return
}

func Register(options ...meerkats.HandlerOption) meerkats.ContextOption {
	return meerkats.ContextOption(func(ctx *meerkats.Context) {
		ctx.Register(New(options...))
	})
}

func (h *stdLogger) SetTimeLayout(layout string) {
	h.TimeLayout = layout
}
func (h *stdLogger) GetTimeLayout() string {
	return h.TimeLayout
}

func (h *stdLogger) SetLevel(level meerkats.Level) {
	h.Level = level
}
func (h *stdLogger) GetLevel() meerkats.Level {
	return h.Level
}
func (h *stdLogger) Dispose() {
	pool.Put(h)
	h.mu.Unlock()
}
func (h *stdLogger) Clone() (meerkats.Handler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	clone := New()
	clone.Level = h.Level
	clone.TimeLayout = h.TimeLayout
	clone.bytes = h.bytes[0:]
	return clone
}

func (h *stdLogger) AddBool(key string, value bool) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteRune('=')
	h.bytes = strconv.AppendBool(buff.Bytes(), value)
}
func (h *stdLogger) AddString(key string, value string) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteString("=\"")
	buff.WriteString(value)
	buff.WriteRune('"')
}
func (h *stdLogger) AddInt(key string, value int) {
	h.AddInt64(key, int64(value))
}
func (h *stdLogger) AddInt64(key string, value int64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteRune('=')
	h.bytes = strconv.AppendInt(h.bytes, value, 10)
}
func (h *stdLogger) AddUint(key string, value uint) {
	h.AddUint64(key, uint64(value))
}
func (h *stdLogger) AddUint64(key string, value uint64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteRune('=')
	h.bytes = strconv.AppendUint(h.bytes, value, 10)
}
func (h *stdLogger) AddFloat32(key string, value float32) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteRune('=')
	h.bytes = strconv.AppendFloat(h.bytes, float64(value), 'f', -1, 32)
}
func (h *stdLogger) AddFloat64(key string, value float64) {
	defer h.mu.Unlock()
	h.mu.Lock()
	buff := bytes.NewBuffer(h.bytes)

	buff.WriteRune(' ')
	buff.WriteString(key)
	buff.WriteRune('=')
	h.bytes = strconv.AppendFloat(h.bytes, value, 'f', -1, 64)
}
func (h *stdLogger) AddObject(key string, value interface{}) {
	h.AddString(key, fmt.Sprintf("%+v", value))
}

func (h *stdLogger) AddFields(fs ...meerkats.KeyValue) {
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

func (h *stdLogger) Log(level meerkats.Level, timestamp time.Time, msg string) {
	h.mu.Lock()
	defer h.Dispose()

	buf := bytes.NewBufferString("\033[")
	buf.WriteString(COLORS[level])
	buf.WriteRune('m')

	buf.WriteString("level=\"")
	buf.WriteString(level.String())

	buf.WriteString("\" timestamp=\"")
	buf.WriteString(timestamp.Format(h.TimeLayout))

	buf.WriteString("\" message=\"")
	buf.WriteString(msg)
	buf.WriteRune('"')

	buf.Write(h.bytes)

	buf.WriteString("\033[0m")

	log.Println(buf.String())
}