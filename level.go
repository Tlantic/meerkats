package meerkats

import "strings"

const (
	LEVEL_TRACE		Level	=	1 << iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_PANIC
)

const LEVEL_ALL = LEVEL_TRACE | LEVEL_DEBUG | LEVEL_INFO | LEVEL_WARNING | LEVEL_ERROR | LEVEL_FATAL | LEVEL_PANIC

type Level uint8

func (l Level) Highest() Level {
	if ( l&LEVEL_PANIC != 0 ) {
		return LEVEL_PANIC
	}
	if ( l&LEVEL_FATAL != 0 ) {
		return LEVEL_FATAL
	}
	if ( l&LEVEL_ERROR != 0 ) {
		return LEVEL_ERROR
	}
	if ( l&LEVEL_WARNING != 0 ) {
		return LEVEL_WARNING
	}
	if ( l&LEVEL_INFO != 0 ) {
		return LEVEL_INFO
	}
	if ( l&LEVEL_DEBUG != 0 ) {
		return LEVEL_DEBUG
	}

	return LEVEL_TRACE
}

func ( l Level ) String() string {


	sl := make([]string, 0, 7)

	if ( l&LEVEL_TRACE != 0 ) {
		sl = append(sl,"TRACE")
	}
	if ( l&LEVEL_DEBUG != 0 ) {
		sl = append(sl,"DEBUG")
	}
	if ( l&LEVEL_INFO != 0 ) {
		sl = append(sl,"INFO")
	}
	if ( l&LEVEL_WARNING != 0) {
		sl = append(sl,"WARNING")
	}
	if ( l&LEVEL_ERROR != 0) {
		sl = append(sl,"ERROR")
	}
	if ( l&LEVEL_FATAL != 0) {
		sl = append(sl,"FATAL")
	}
	if ( l&LEVEL_PANIC != 0) {
		sl = append(sl,"PANIC")
	}

	return strings.Join(sl, "|")
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}