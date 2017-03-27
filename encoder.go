package meerkats

type Encoder interface {
	AddBool(key string, value bool)
	AddString(key string, value string)
	AddInt(key string, value int)
	AddInt64(key string, value int64)
	AddUint(key string, value uint)
	AddUint64(key string, value uint64)
	AddFloat32(key string, value float32)
	AddFloat64(key string, value float64)
	AddJSON(key string, value interface{})
	AddError(err error)
	Add(key string, value interface{})
	With(fields ...Field)
}
