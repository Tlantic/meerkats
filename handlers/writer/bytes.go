package writer

import "strconv"

func appendBool(key string, value bool) []byte {
	return strconv.AppendBool(append(append([]byte(key), '=')), value)
}

func appendString(key string, value string) []byte {
	return append(append(append([]byte(key), []byte("=\"")...), value...), '"')
}

func appendInt64(key string, value int64) []byte {
	return strconv.AppendInt(append(append([]byte(key), '=')), value, 10)
}

func appendUint64(key string, value uint64) []byte {
	return strconv.AppendUint(append(append([]byte(key), '=')), value, 10)
}

func appendFloat32(key string, value float32) []byte {
	return strconv.AppendFloat(append(append([]byte(key), '=')), float64(value), 'f', -1, 32)
}

func appendFloat64(key string, value float64) []byte {
	return strconv.AppendFloat(append(append([]byte(key), '=')), value, 'f', -1, 64)
}
