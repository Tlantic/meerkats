package meerkats

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"strconv"
)

type FieldSet []Field

// Deprecated: Use log.Field
type Field struct {
	Key            string
	Type           FieldType
	ValueBool      bool
	ValueString    string
	ValueInt64     int64
	ValueUint64    uint64
	ValueFloat32   float32
	ValueFloat64   float64
	ValueInterface interface{}
}

func (v Field) Apply(l Logger) {
	l.With(v)
}

func (v Field) GetType() FieldType {
	return v.Type
}

func (v Field) Get() interface{} {
	switch v.Type {
	case TypeBool:
		return v.ValueBool
	case TypeString:
		return v.ValueString
	case TypeInt64:
		return v.ValueInt64
	case TypeUint64:
		return v.ValueUint64
	case TypeFloat32:
		return v.ValueFloat32
	case TypeFloat64:
		return v.ValueFloat64
	default:
		return v.ValueInterface
	}
}
func (v *Field) Set(value interface{}) {
	switch value := value.(type) {
	case string:
		v.Type = TypeString
		v.ValueString = value
	case bool:
		v.Type = TypeBool
		v.ValueBool = value
	case int:
		v.Type = TypeInt64
		v.ValueInt64 = int64(value)
	case int8:
		v.Type = TypeInt64
		v.ValueInt64 = int64(value)
	case int16:
		v.Type = TypeInt64
		v.ValueInt64 = int64(value)
	case int32:
		v.Type = TypeInt64
		v.ValueInt64 = int64(value)
	case int64:
		v.Type = TypeInt64
		v.ValueInt64 = value
	case uint:
		v.Type = TypeUint64
		v.ValueUint64 = uint64(value)
	case uint8:
		v.Type = TypeUint64
		v.ValueUint64 = uint64(value)
	case uint16:
		v.Type = TypeUint64
		v.ValueUint64 = uint64(value)
	case uint32:
		v.Type = TypeUint64
		v.ValueUint64 = uint64(value)
	case uint64:
		v.Type = TypeUint64
		v.ValueUint64 = value
	case float32:
		v.Type = TypeFloat32
		v.ValueFloat32 = value
	case float64:
		v.Type = TypeFloat64
		v.ValueFloat64 = value
	case error:
		v.Type = TypeError
		v.ValueInterface = value
	default:
		v.Type = TypeInterface
		v.ValueInterface = value
	}

}
func (v Field) String() string {
	switch v.Type {
	case TypeString:
		return v.ValueString
	case TypeBool:
		return strconv.FormatBool(v.ValueBool)
	case TypeInt64:
		return strconv.FormatInt(v.ValueInt64, 10)
	case TypeUint64:
		return strconv.FormatUint(v.ValueUint64, 10)
	case TypeFloat32:
		return strconv.FormatFloat(float64(v.ValueFloat32), 'f', -1, 32)
	case TypeFloat64:
		return strconv.FormatFloat(v.ValueFloat64, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v.ValueInterface)
	}
}

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
