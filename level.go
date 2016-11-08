package meerkats


type Level uint8

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

var levels = [...]string {
	LEVEL_TRACE: "TRACE",
	LEVEL_DEBUG: "DEBUG",
	LEVEL_INFO: "INFO",
	LEVEL_WARNING: "WARNING",
	LEVEL_ERROR: "ERROR",
	LEVEL_FATAL: "FATAL",
	LEVEL_PANIC: "PANIC",
}


func ( l Level ) String() string {


	if s := levels[l]; s != "" {
		return s
	}

	return "UNKNOWN"
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}