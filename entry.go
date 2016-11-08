package meerkats

import (
	"time"
	"fmt"
)

const FIELD_STRUCT_TAG = "log"

var _ ILogger = (*Entry)(nil)


type Entry struct {
	meerkat   	*Meerkat		`json:"-" xml:"-"`
	Timestamp 	time.Time		`json:"timestamp" xml:"timestamp" mapstructure:"timestamp"`
	TimeLayout	string			`json:"-" xml:"-" mapstructure:"timeLayout"`
	TraceId  	string			`json:"id" xml:"id" mapstructure:"id"`
	Level     	Level			`json:"level" xml:"level" mapstructure:"level"`


	Message   	string			`json:"message" xml:"message" mapstructure:"message"`
	Fields    	Fields			`json:"fields" xml:"fields" mapstructure:"fields"`
}

func NewEntry(m *Meerkat) *Entry {
	return &Entry{
		meerkat: m,
		Timestamp: time.Now(),
		TimeLayout: m.TimeLayout,
		Fields: Fields{},
	}
}

func (e *Entry) String() string {
	return fmt.Sprintf("timestamp=\"%s\" level=\"%s\" message=\"%s\" %s",
		e.Timestamp.Format(e.TimeLayout), e.Level, e.Message, e.Fields.String())
}

func ( e *Entry ) With(fields ... interface{}) ILogger {
	e.Fields.Merge(fields...)
	return e
}
func ( e *Entry ) WithField( name string, value interface{}) ILogger {
	e.Fields[name] = value
	return e
}
func ( e *Entry ) WithFields(fields ... interface{}) ILogger {
	e.Fields.Merge(fields...)
	return e
}



func ( e *Entry ) Log(level interface{}, a ... interface{}) {
	e.meerkat.wg.Add(1)
	e.Level = level.(Level)
	e.Message = fmt.Sprint(a...)
	e.meerkat.queue <- *e
}
func ( e *Entry ) Logln(level interface{}, a ... interface{}) {
	e.meerkat.wg.Add(1)
	e.Level = level.(Level)
	e.Message = fmt.Sprint(a...)
	e.meerkat.queue <- *e
}
func ( e *Entry ) Logf(level interface{}, format string, v ... interface{}) {
	e.meerkat.wg.Add(1)
	e.Level = level.(Level)
	e.Message = fmt.Sprintf(format, v...)
	e.meerkat.queue <- *e
}

func ( e *Entry ) Print(a ... interface{}) {
	e.Log(LEVEL_TRACE, a...)
}
func ( e *Entry ) Println(a ... interface{}) {
	e.Logln(LEVEL_TRACE, a...)
}
func ( e *Entry ) Printf(format string, v ... interface{}) {
	e.Logf(LEVEL_TRACE, format, v...)
}

func ( e *Entry ) Trace(a ... interface{}) {
	e.Log(LEVEL_TRACE, a...)
}
func ( e *Entry ) Traceln(a ... interface{}) {
	e.Logln(LEVEL_TRACE, a...)
}
func ( e *Entry ) Tracef(format string, v ... interface{}) {
	e.Logf(LEVEL_TRACE, format, v...)
}

func ( e *Entry ) Debug(a ... interface{}) {
	e.Log(LEVEL_DEBUG, a...)
}
func ( e *Entry ) Debugln(a ... interface{}) {
	e.Logln(LEVEL_DEBUG, a...)
}
func ( e *Entry ) Debugf(format string, v ... interface{}) {
	e.Logf(LEVEL_DEBUG, format, v...)
}

func ( e *Entry ) Info(a ... interface{}) {
	e.Log(LEVEL_INFO, a...)
}
func ( e *Entry ) Infoln(a ... interface{}) {
	e.Logln(LEVEL_INFO, a...)
}
func ( e *Entry ) Infof(format string, v ... interface{}) {
	e.Logf(LEVEL_INFO, format, v...)
}

func ( e *Entry ) Warning(a ... interface{}) {
	e.Log(LEVEL_WARNING, a...)
}
func ( e *Entry ) Warningln(a ... interface{}) {
	e.Logln(LEVEL_WARNING, a...)
}
func ( e *Entry ) Warningf(format string, v ... interface{}) {
	e.Logf(LEVEL_WARNING, format, v...)
}

func ( e *Entry ) Error(a ... interface{}) {
	e.Log(LEVEL_ERROR, a...)
}
func ( e *Entry ) Errorln(a ... interface{}) {
	e.Logln(LEVEL_ERROR, a...)
}
func ( e *Entry ) Errorf(format string, v ... interface{}) {
	e.Logf(LEVEL_ERROR, format, v...)
}

func ( e *Entry ) Fatal(a ... interface{}) {
	e.Log(LEVEL_FATAL, a...)
}
func ( e *Entry ) Fatalln(a ... interface{}) {
	e.Logln(LEVEL_FATAL, a...)
}
func ( e *Entry ) Fatalf(format string, v ... interface{}) {
	e.Logf(LEVEL_FATAL, format, v...)
}

func ( e *Entry ) Panic(a ... interface{}) {
	e.Log(LEVEL_PANIC, a...)
}
func ( e *Entry ) Panicln(a ... interface{}) {
	e.Logln(LEVEL_PANIC, a...)
}
func ( e *Entry ) Panicf(format string, v ... interface{}) {
	e.Logf(LEVEL_PANIC, format, v...)
}