package meerkats

import "fmt"

type FieldSet []KeyValue

type KeyValue struct {
	Key            string
	Type           FieldType
	ValueBool      bool
	ValueString    string
	ValueInt       int64
	ValueUint      uint64
	ValueFloat32   float32
	ValueFloat64   float64
	ValueInterface interface{}
}

func Field(key string, value interface{}) (f KeyValue) {
	f = KeyValue{Key: key}
	f.Set(value)
	return
}
func String(key string, value string) KeyValue {
	return KeyValue{Key: key, Type: TypeString, ValueString: value}
}
func Bool(key string, value bool) KeyValue {
	return KeyValue{Key: key, Type: TypeBool, ValueBool: value}
}
func Int(key string, value int) KeyValue {
	return KeyValue{Key: key, Type: TypeInt, ValueInt: int64(value)}
}
func Int64(key string, value int64) KeyValue {
	return KeyValue{Key: key, Type: TypeInt, ValueInt: value}
}
func Uint(key string, value uint) KeyValue {
	return KeyValue{Key: key, Type: TypeUint, ValueUint: uint64(value)}
}
func Uint64(key string, value uint64) KeyValue {
	return KeyValue{Key: key, Type: TypeUint, ValueUint: value}
}
func Float32(key string, value float32) KeyValue {
	return KeyValue{Key: key, Type: TypeFloat32, ValueFloat32: value}
}
func Float64(key string, value float64) KeyValue {
	return KeyValue{Key: key, Type: TypeFloat64, ValueFloat64: value}
}

func (v *KeyValue) clear() {
	v.ValueInterface	= nil
	v.ValueString		= ""
	v.ValueBool			= false
	v.ValueInt			= 0
	v.ValueUint			= 0
	v.ValueFloat32		= 0
	v.ValueFloat64		= 0
}

func (v *KeyValue) GetType() FieldType {
	return v.Type
}

func (v *KeyValue) Get() interface{} {
	switch v.Type {
	case TypeString:
		return v.ValueString
	case TypeBool:
		return v.ValueBool
	case TypeInt64:
		return v.ValueInt
	case TypeUint64:
		return v.ValueUint
	case TypeFloat32:
		return v.ValueFloat32
	case TypeFloat64:
		return v.ValueFloat64
	default:
		return v.ValueInterface
	}
}
func (v *KeyValue) Set(value interface{}) {
	switch typedValue := value.(type) {
	case string:
		v.SetString(typedValue)
	case bool:
		v.SetBool(typedValue)
	case int:
		v.SetInt(typedValue)
	case int8:
		v.SetInt64(int64(typedValue))
	case int16:
		v.SetInt64(int64(typedValue))
	case int32:
		v.SetInt64(int64(typedValue))
	case int64:
		v.SetInt64(typedValue)
	case uint:
		v.SetUint(typedValue)
	case uint8:
		v.SetUint64(uint64(typedValue))
	case uint16:
		v.SetUint64(uint64(typedValue))
	case uint32:
		v.SetUint64(uint64(typedValue))
	case uint64:
		v.SetUint64(typedValue)
	case float32:
		v.SetFloat32(typedValue)
	case float64:
		v.SetFloat64(typedValue)
	default:
		v.SetInterface(typedValue)
	}
}

func (v *KeyValue) GetString() string {
	return v.ValueString
}
func (v *KeyValue) SetString(value string) {
	v.Type = TypeString
	v.ValueString = value
}

func (v *KeyValue) GetBool() bool {
	return v.ValueBool
}
func (v *KeyValue) SetBool(value bool) {
	v.Type = TypeBool
	v.ValueBool = value
}

func (v *KeyValue) GetInt() int {
	return int(v.ValueInt)
}
func (v *KeyValue) SetInt(value int) {
	v.Type = TypeInt64
	v.ValueInt = int64(value)
}

func (v *KeyValue) GetInt64() int64 {
	return v.ValueInt
}
func (v *KeyValue) SetInt64(value int64) {
	v.Type = TypeInt64
	v.ValueInt = value
}

func (v *KeyValue) GetUint() uint {
	return uint(v.ValueUint)
}
func (v *KeyValue) SetUint(value uint) {
	v.Type = TypeUint64
	v.ValueUint = uint64(value)
}

func (v *KeyValue) GetUint64() uint64 {
	return v.ValueUint
}
func (v *KeyValue) SetUint64(value uint64) {
	v.Type = TypeUint64
	v.ValueUint = value
}

func (v *KeyValue) GetFloat32() float32 {
	return v.ValueFloat32
}
func (v *KeyValue) SetFloat32(value float32) {
	v.Type = TypeFloat32
	v.ValueFloat32 = value
}

func (v *KeyValue) GetFloat64() float64 {
	return v.ValueFloat64
}
func (v *KeyValue) SetFloat64(value float64) {
	v.Type = TypeFloat64
	v.ValueFloat64 = value
}

func (v *KeyValue) GetInterface() interface{} {
	return v.ValueInterface
}
func (v *KeyValue) SetInterface(value interface{}) {
	v.Type = TypeInterface
	v.ValueInterface = value
}

func (v *KeyValue) String() string {
	if ( v.Type == TypeString ) {
		return v.GetString()
	}
	return fmt.Sprint(v.Get())
}
