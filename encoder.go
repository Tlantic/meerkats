package meerkats

import "github.com/opentracing/opentracing-go/log"

type Encoder interface {
	log.Encoder
	EmitUint(key string, value uint)
	EmitJSON(key string, value interface{})
	EmitError(err error)
	EmitField(fields ...log.Field)
}
