package meerkats

import (
	"time"
	"sync"
)

type Callback func()
type EntryQueue chan Entry
type EntryHandler func(Entry, Callback)

type MeerkatOptions struct {
	Level      		Level
	TimeLayout 		string
	MaxWorkers 		int
}
type Meerkat struct {
	Level           Level
	TimeLayout      string

	wg              sync.WaitGroup
	entryQueue      EntryQueue
	workerPool      *workperPool
	handlersCatalog *catalog
	closeChan       chan bool
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
		Level: 				opts.Level,
		TimeLayout: 		timeLayout,

		wg:       			sync.WaitGroup{},
		entryQueue:			make(EntryQueue),
		workerPool:			newWorkerPool(opts.MaxWorkers),
		handlersCatalog:	newCatalog(),
		closeChan: 			make(chan bool),
	}

	for i := 0; i < maxWorkers; i++ {
		m.AddWorker()
	}

	go func() {
		defer close(m.entryQueue)

		for {
			select {
			case entry := <-m.entryQueue:
				go func(entry Entry) {
					queue := <-m.workerPool.queueChans
					queue <- entry
				}(entry)
			case <-m.closeChan:
				m.workerPool.Close()
				return
			}

		}
	}()

	return m
}
func ( m *Meerkat ) AddWorker() {

	go func() {
		queue := make(EntryQueue)
		defer close(queue)

		closeChan := make(chan bool)
		m.workerPool.AddCloseChan(closeChan)

		for {
			m.workerPool.queueChans <- queue
			select {
			case entry := <-queue:
				if ( entry.Level >= m.Level ) {
					handlers := m.handlersCatalog.handlers[entry.Level]
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
	m.closeChan <- true
	m.wg.Wait()
}
func ( m *Meerkat ) RegisterHandler(level Level, handlers ... EntryHandler) {
	m.handlersCatalog.Add(level, handlers...)
}


