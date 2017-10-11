package meerkats

import "github.com/opentracing/opentracing-go/log"

var root = New(TRACE)

func SetLevel(lvl Level) {
	root.SetLevel(lvl)
}

func Register(hs ...Handler) {
	root.Register(hs...)
}

func SetTag(key string, value interface{}) {
	root.SetTag(key, value)
}
func SetMeta(key string, value string) {
	root.SetTag(key, value)
}

func GetTag(key string) interface{} {
	return root.GetMeta(key)
}
func GetMeta(key string) string {
	return root.GetMeta(key)
}

// Deprecated: Use EmitBool
func AddBool(key string, value bool) {
	EmitBool(key, value)
}
func EmitBool(key string, value bool) {
	root.EmitBool(key, value)
}

// Deprecated: Use EmitString
func AddString(key string, value string) {
	EmitString(key, value)
}
func EmitString(key string, value string) {
	root.EmitString(key, value)
}

// Deprecated: Use EmitInt
func AddInt(key string, value int) {
	EmitInt(key, value)
}
func EmitInt(key string, value int) {
	root.EmitInt(key, value)
}

func EmitInt32(key string, value int32) {
	root.EmitInt32(key, value)
}

// Deprecated: Use EmitIn64
func AddInt64(key string, value int64) {
	EmitInt64(key, value)
}
func EmitInt64(key string, value int64) {
	root.EmitInt64(key, value)
}

// Deprecated: Use EmitUint
func AddUint(key string, value uint) {
	EmitUint(key, value)
}
func EmitUint(key string, value uint) {
	root.EmitUint(key, value)
}

// Deprecated: Use EmitUint64
func AddUint64(key string, value uint64) {
	EmitUint64(key, value)
}
func EmitUint64(key string, value uint64) {
	root.EmitUint64(key, value)
}

// Deprecated: Use EmitFloat32
func AddFloat32(key string, value float32) {
	EmitFloat32(key, value)
}
func EmitFloat32(key string, value float32) {
	root.EmitFloat32(key, value)
}

// Deprecated: Use EmitFloat64
func AddFloat64(key string, value float64) {
	EmitFloat64(key, value)
}
func EmitFloat64(key string, value float64) {
	root.EmitFloat64(key, value)
}

// Deprecated: Use EmitJSON
func AddJSON(key string, value interface{}) {
	EmitJSON(key, value)
}
func EmitJSON(key string, value interface{}) {
	root.EmitJSON(key, value)
}

// Deprecated: Use EmitError
func AddError(err error) {
	EmitError(err)
}
func EmitError(err error) {
	root.EmitError(err)
}

// Deprecated: Use EmitObject
func AddObject(key string, value interface{}) {
	EmitObject(key, value)
}
func EmitObject(key string, value interface{}) {
	root.EmitObject(key, value)
}

// Deprecated: Use EmitField
func With(fields ...Field) {
	root.With(fields...)
}
func EmitField(fields ...log.Field) {
	root.EmitField(fields...)
}

func Log(level Level, msg string, fields ...log.Field) {
	root.Log(level, msg, fields...)
}

func Trace(msg string, fields ...log.Field) {
	root.Trace(msg, fields...)
}

func Debug(msg string, fields ...log.Field) {
	root.Debug(msg, fields...)
}
func Info(msg string, fields ...log.Field) {
	root.Info(msg, fields...)
}
func Warn(msg string, fields ...log.Field) {
	root.Warn(msg, fields...)
}
func Error(msg string, fields ...log.Field) {
	root.Error(msg, fields...)
}
func Panic(msg string, fields ...log.Field) {
	root.Panic(msg, fields...)
}
func Fatal(msg string, fields ...log.Field) {
	root.Fatal(msg, fields...)
}

func Clone() Logger {
	return root.Clone()
}
