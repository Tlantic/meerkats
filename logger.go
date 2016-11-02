package meerkats


type ILogger interface {
	Print()
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal()
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic()
	Panicf(string, ...interface{})
	Panicln(...interface{})
}