package meerkats

import (
	"time"
	"sync"
)

type Callback func()
type EntryQueue chan Entry
type EntryHandler func(Entry, Callback)

type MeerkatOptions struct {
	Level      	Level
	TimeLayout	string
	MaxWorkers 	uint
	QueueSize	uint
}
type Meerkat struct {
	Level      Level
	TimeLayout string

	wg         sync.WaitGroup
	queue      EntryQueue
	manager    *workerManager
	handlers   *catalog
	closeChan  chan bool
}

func NewMeerkat(opts MeerkatOptions) *Meerkat {

	var maxWorkers 	uint 	= 1
	var queueSize	uint 	= 200
	var timeLayout 	string 	= time.RFC3339

	if ( opts.TimeLayout != "" ) {
		timeLayout = opts.TimeLayout
	}

	if (opts.MaxWorkers > 0) {
		maxWorkers = opts.MaxWorkers
	}

	if (opts.QueueSize > 0) {
		queueSize = opts.QueueSize
	}

	m := &Meerkat{
		Level: opts.Level,
		TimeLayout: timeLayout,

		wg: sync.WaitGroup{},
		queue: make(EntryQueue, queueSize),
		manager: newWorkerManager(opts.MaxWorkers),
		handlers: newHandlerManager(),
		closeChan: make(chan bool),
	}

	var i uint
	for ; i < maxWorkers; i++ {
		m.AddWorker()
	}

	go func() {
		defer close(m.queue)

		for {
			select {
			case entry := <-m.queue:
				go func(entry Entry) {
					queue := <-m.manager.pool
					queue <- entry
				}(entry)
			case <-m.closeChan:
				m.manager.Close()
				return
			}

		}
	}()

	return m
}

//noinspection GoUnusedExportedFunction
func New(opts *MeerkatOptions) *Meerkat {
	if ( opts == nil) {
		opts = &MeerkatOptions{}
	}
	return NewMeerkat(*opts)
}

func ( m *Meerkat ) AddWorker() {

	go func() {
		queue := make(EntryQueue)
		defer close(queue)

		closeChan := make(chan bool)
		m.manager.AddCloseChan(closeChan)

		for {
			m.manager.pool <- queue
			select {
			case entry := <-queue:
				if ( entry.Level >= m.Level ) {
					handlers := m.handlers.handlers[entry.Level]
					if ( handlers != nil ) {
						m.wg.Add(len(handlers))
						for _, h := range handlers {
							go h(entry, m.wg.Done)
						}
					}
				}
				m.wg.Done()
			case <-closeChan:
				return
			}

		}
	}()
}
func ( m *Meerkat ) Close() {
	m.wg.Wait()
	m.closeChan <- true
}
func ( m *Meerkat ) RegisterHandler(level Level, handlers ... EntryHandler) {
	m.handlers.Add(level, handlers...)
}


