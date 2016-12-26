package meerkats

import "sync"

var mu = sync.Mutex{}
var root = New(TRACE)

func SetLevel(lvl Level) {
	root.SetLevel(lvl)
}

func Register(hs ... Handler) {
	root.Register(hs...)
}

func SetMeta(key string, value string) {
	root.SetMeta(key, value)
}
func GetMeta(key string) string {
	return root.GetMeta(key)
}

func AddBool(key string, value bool) {
	root.AddBool(key, value)
}
func AddString(key string, value string) {
	root.AddString(key, value)
}
func AddInt(key string, value int) {
	root.AddInt(key, value)
}
func AddIn64(key string, value int64) {
	root.AddInt64(key, value)
}
func AddUint(key string, value uint) {
	root.AddUint(key, value)
}
func AddUint64(key string, value uint64) {
	root.AddUint64(key, value)
}
func AddFloat32(key string, value float32) {
	root.AddFloat32(key, value)
}
func AddFloat64(key string, value float64) {
	root.AddFloat64(key, value)
}
func AddObject(key string, value interface{}) {
	root.Add(key, value)
}
func With(fields ...Field) {
	root.With(fields...)
}


func Log(level Level, msg string, fields ...Field) {
	root.Log(level, msg, fields...)
}


func Trace(msg string, fields ...Field) {
	root.Trace(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	root.Debug(msg, fields...)
}
func Info(msg string, fields ...Field) {
	root.Info(msg, fields...)
}
func Warn(msg string, fields ...Field) {
	root.Warn(msg, fields...)
}
func Error(msg string, fields ...Field) {
	root.Error(msg, fields...)
}
func Panic(msg string, fields ...Field) {
	root.Panic(msg, fields...)
}
func Fatal(msg string, fields ...Field) {
	root.Fatal(msg, fields...)
}