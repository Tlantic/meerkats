package meerkats


type ILogger interface {
	With(... interface{}) ILogger
	WithField(string, interface{}) ILogger
	WithFields(... interface{}) ILogger

	Log(interface{}, ... interface{})
	Logln(interface{}, ... interface{})
	Logf(interface{}, string, ... interface{})

	Print(... interface{})
	Println(... interface{})
	Printf(string, ... interface{})

	Trace(... interface{})
	Traceln(... interface{})
	Tracef(string, ... interface{})

	Debug(... interface{})
	Debugln(... interface{})
	Debugf(string, ... interface{})

	Info(... interface{})
	Infoln(... interface{})
	Infof(string, ... interface{})

	Warning(... interface{})
	Warningln(... interface{})
	Warningf(string, ... interface{})

	Error(... interface{})
	Errorln(... interface{})
	Errorf(string, ... interface{})

	Fatal(... interface{})
	Fatalln(... interface{})
	Fatalf(string, ... interface{})

	Panic(... interface{})
	Panicln(... interface{})
	Panicf(string, ... interface{})
	String() string
}