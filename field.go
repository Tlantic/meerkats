package meerkats

import (
	"encoding/json"
	"github.com/opentracing/opentracing-go/log"
)

func NewField(key string, value interface{}) (f log.Field) {
	return log.Object(key, value)
}
func String(key string, value string) log.Field {
	return log.String(key, value)
}
func Bool(key string, value bool) log.Field {
	return log.Bool(key, value)
}
func Int(key string, value int) log.Field {
	return log.Int(key, value)
}
func Int32(key string, value int32) log.Field {
	return log.Int32(key, value)
}
func Int64(key string, value int64) log.Field {
	return log.Int64(key, value)
}

func Uint(key string, value uint) log.Field {
	return log.Uint64(key, uint64(value))
}
func Uint32(key string, value uint32) log.Field {
	return log.Uint32(key, value)
}
func Uint64(key string, value uint64) log.Field {
	return log.Uint64(key, value)
}

func Float32(key string, value float32) log.Field {
	return log.Float32(key, value)
}
func Float64(key string, value float64) log.Field {
	return log.Float64(key, value)
}
func Object(key string, value interface{}) log.Field {
	return log.Object(key, value)
}
func JSON(key string, value interface{}) log.Field {
	s, _ := json.Marshal(value)
	return log.String(key, string(s))
}
func ErrorString(err error) log.Field {
	return log.Error(err)
}
