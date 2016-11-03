package meerkats

import (
	"time"
	"sync"
)

type Callback func()
type EntryQueue chan Entry
type EntryHandler func(Entry, Callback)


type MeerkatOptions struct {
	Level      Level
	TimeLayout string
	MaxWorkers int
}
type Meerkat struct {
	Level           Level
	TimeLayout      string

	wg              sync.WaitGroup
	entryQueue      EntryQueue
	workers         *WorkerPool
	handlersCatalog *HandlerCatalog
	close           chan bool
}

func NewMeerkat(opts MeerkatOptions) *Meerkat {

	maxWorkers := 1
	timeLayout := time.RFC3339

	if ( opts.TimeLayout != "" ) {
		timeLayout = opts.TimeLayout
	}

	if (opts.MaxWorkers > 0) {
		maxWorkers = opts.MaxWorkers
	}



	m := &Meerkat{
		Level: opts.Level,
		TimeLayout: timeLayout,
		wg:       sync.WaitGroup{},
		entryQueue: make(EntryQueue),
		workers: NewWorkerPool(opts.MaxWorkers),
		handlersCatalog: NewHandlerCatalog(),
		close: make(chan bool),
	}

	for i := 0; i < maxWorkers; i++ {
		m.AddWorker()
	}

	go func() {
		for {
			select {
			case entry := <-m.entryQueue:
				go func(entry Entry) {
					queue := <-m.workers.availableQueues
					queue <- entry
				}(entry)
			case <-m.close:
				m.workers.Close()
				return
			}
		}
	}()

	return m
}
func (m *Meerkat) AddWorker() {
	dispose :=  make(chan bool)
	m.workers.Add( dispose )
	go func(dispose chan bool) {
		queue := make(EntryQueue)
		for {
			m.workers.availableQueues <- queue

			select {
			case entry := <- queue:
				if ( entry.Level >= m.Level ) {
					handlers := m.handlersCatalog.handlers[entry.Level]
					if ( handlers != nil ) {
						m.wg.Add(len(handlers))
						for _, h := range handlers {
							h( entry, m.wg.Done )
						}
					}
				}
				m.wg.Done()
			case <- dispose:
				close(queue)
				return
			}
		}
	}(dispose)
}
func ( m *Meerkat ) Close() {
	m.wg.Wait()
	close(m.entryQueue)
	m.close <- true
}
func (m *Meerkat) RegisterHandler(level Level, handlers ... EntryHandler) {
	m.handlersCatalog.Add(level, handlers...)
}





