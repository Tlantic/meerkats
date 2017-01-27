package writer

import "strconv"

func appendBool(content []byte, key string, value bool) []byte {
	return strconv.AppendBool(append(append(append(content, ' '), key...), '='), value)
}

func appendString(content []byte, key string, value string) []byte {
	return append(append(append(append(append(content, ' '), key...), []byte("=\"")...), value...), '"')
}

func appendInt64(content []byte, key string, value int64) []byte {
	return strconv.AppendInt(append(append(append(append(content, ' '), key...), '=')), value, 10)
}

func appendUint64(content []byte, key string, value uint64) []byte {
	return strconv.AppendUint(append(append(append(content, ' '), key...), '='), value, 10)
}

func appendFloat32(content []byte, key string, value float32) []byte {
	return strconv.AppendFloat(append(append(append(append(content, ' '), key...), '=')), float64(value), 'f', -1, 32)
}

func appendFloat64(content []byte, key string, value float64) []byte {
	return strconv.AppendFloat(append(append(append(append(content, ' '), key...), '=')), value, 'f', -1, 64)
}