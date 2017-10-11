package meerkats

import "github.com/opentracing/opentracing-go/log"

type Encoder interface {
	log.Encoder
	// Deprecated: use EmitString
	AddString(key string, value string)

	// Deprecated: use EmitBool
	AddBool(key string, value bool)

	// Deprecated: use EmitInt
	AddInt(key string, value int)

	// Deprecated: use EmitInt64
	AddInt64(key string, value int64)

	// Deprecated: use AddUint
	AddUint(key string, value uint)
	EmitUint(key string, value uint)

	// Deprecated: use EmitUint64
	AddUint64(key string, value uint64)

	// Deprecated: use EmitFloat32
	AddFloat32(key string, value float32)

	// Deprecated: use EmitFloat64
	AddFloat64(key string, value float64)

	// Deprecated: use EmitObject
	Add(key string, value interface{})

	// Deprecated: use EmitJSON
	AddJSON(key string, value interface{})
	EmitJSON(key string, value interface{})

	// Deprecated: use EmitError
	AddError(err error)
	EmitError(err error)

	// Deprecated: use EmitField
	With(fields ...Field)
	EmitField(fields ...log.Field)
}
