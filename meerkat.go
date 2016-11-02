package meerkats

import (
	"time"
	"sync"
)

type WorkerPool chan chan Entry
type EntryQueue chan Entry
type EntryHandler func(Entry)

type MeerkatOptions struct {
	Level      Level
	TimeLayout string
	Dispatcher IEntryDispatcher
	MaxWorkers int
}
type Meerkat struct {
	Level      Level
	TimeLayout string
	Dispatcher IEntryDispatcher
	wg         sync.WaitGroup
	queue      EntryQueue
	close      chan bool
}

func NewMeerkat(opts MeerkatOptions) *Meerkat {

	var dispatcher IEntryDispatcher

	maxWorkers := 1
	timeLayout := time.RFC3339


	if ( opts.TimeLayout != "" ) {
		timeLayout = opts.TimeLayout
	}

	if (opts.Dispatcher != nil) {
		dispatcher = opts.Dispatcher
	} else {
		dispatcher = NewEntryDispatcher()
	}

	if (opts.MaxWorkers > 0) {
		maxWorkers = opts.MaxWorkers
	}

	m := &Meerkat{
		TimeLayout: timeLayout,
		Dispatcher: dispatcher,
		wg:       sync.WaitGroup{},
		queue: make(EntryQueue),
		close: make(chan bool),
	}

	pool := make(WorkerPool, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		m.Dispatcher.NewWorker(opts.Level, pool, m.wg.Done)
	}

	go func() {
		for {
			select {
			case entry := <-m.queue:
				go func(entry Entry) {
					worker := <-pool
					worker <- entry
				}(entry)
			case <-m.close:
				close(m.queue)
				m.Dispatcher.Close()
				return
			}
		}
	}()

	return m
}
func ( m *Meerkat ) Close() {
	m.wg.Wait()
	m.close <- true
}
func (m *Meerkat) RegisterHandler(level Level, handlers ... EntryHandler) {
	m.Dispatcher.Register(level, handlers...)
}

func ( m *Meerkat) WithField(name string, value interface{}) *Entry {
	return NewEntry(m).WithField(name, value)
}
func ( m *Meerkat) WithFields(fields ... Fields) *Entry {
	return NewEntry(m).WithFields(fields...)
}

func ( m *Meerkat ) Log(level Level, a ... interface{}) {
	NewEntry(m).Log(level, a...)
}
func ( m *Meerkat ) Logln(level Level, a ... interface{}) {
	NewEntry(m).Logln(level, a...)
}
func ( m *Meerkat ) Logf(level Level, format string, v ... interface{}) {
	NewEntry(m).Logf(level, format, v...)
}

func ( m *Meerkat ) Print(a ... interface{}) {
	NewEntry(m).Print(LEVEL_TRACE, a...)
}
func ( m *Meerkat ) Println(a ... interface{}) {
	NewEntry(m).Println(LEVEL_TRACE, a...)
}
func ( m *Meerkat ) Printf(format string, v ... interface{}) {
	NewEntry(m).Printf(LEVEL_TRACE, format, v...)
}

func ( m *Meerkat ) Trace(a ... interface{}) {
	NewEntry(m).Trace(a...)
}
func ( m *Meerkat ) Traceln(a ... interface{}) {
	NewEntry(m).Traceln(a...)
}
func ( m *Meerkat ) Tracef(format string, v ... interface{}) {
	NewEntry(m).Tracef(format, v...)
}

func ( m *Meerkat ) Debug(a ... interface{}) {
	NewEntry(m).Debug(a...)
}
func ( m *Meerkat ) Debugln(a ... interface{}) {
	NewEntry(m).Debugln(a...)
}
func ( m *Meerkat ) Debugf(format string, v ... interface{}) {
	NewEntry(m).Debugf(format, v...)
}

func ( m *Meerkat ) Info(a ... interface{}) {
	NewEntry(m).Info(a...)
}
func ( m *Meerkat ) Infoln(a ... interface{}) {
	NewEntry(m).Infoln(a...)
}
func ( m *Meerkat ) Infof(format string, v ... interface{}) {
	NewEntry(m).Infof(format, v...)
}

func ( m *Meerkat ) Warning(a ... interface{}) {
	NewEntry(m).Warning(a...)
}
func ( m *Meerkat ) Warningln(a ... interface{}) {
	NewEntry(m).Warningln(a...)
}
func ( m *Meerkat ) Warningf(format string, v ... interface{}) {
	NewEntry(m).Warningf(format, v...)
}

func ( m *Meerkat ) Error(a ... interface{}) {
	NewEntry(m).Error(a...)
}
func ( m *Meerkat ) Errorln(a ... interface{}) {
	NewEntry(m).Errorln(a...)
}
func ( m *Meerkat ) Errorf(format string, v ... interface{}) {
	NewEntry(m).Errorf(format, v...)
}

func ( m *Meerkat ) Fatal(a ... interface{}) {
	NewEntry(m).Fatal(a...)
}
func ( m *Meerkat ) Fatalln(a ... interface{}) {
	NewEntry(m).Fatalln(a...)
}
func ( m *Meerkat ) Fatalf(format string, v ... interface{}) {
	NewEntry(m).Fatalf(format, v...)
}

func ( m *Meerkat ) Panic(a ... interface{}) {
	NewEntry(m).Panic(a...)
}
func ( m *Meerkat ) Panicln(a ... interface{}) {
	NewEntry(m).Panicln(a...)
}
func ( m *Meerkat ) Panicf(format string, v ... interface{}) {
	NewEntry(m).Panicf(format, v...)
}

