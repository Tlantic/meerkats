package meerkats

type Level uint8

const (
	TRACE Level	=	1 << iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)


const LEVEL_ALL = TRACE | DEBUG | INFO | WARNING | ERROR | FATAL | PANIC


var Levels = [...]string {
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	INFO: "INFO",
	WARNING: "WARNING",
	ERROR: "ERROR",
	FATAL: "FATAL",
	PANIC: "PANIC",
}

func ( l Level ) String() string {
	return Levels[l]
}
