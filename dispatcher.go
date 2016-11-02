package meerkats

import (
	"sync"
)

type disposeChannels struct {
	mu			sync.Mutex
	channels  	[]chan bool
}
func (d *disposeChannels) Add( ch chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.channels = append( d.channels, ch )
}


type handlerCatalog struct {
	mu			sync.Mutex
	handlers  	map[Level][]EntryHandler
}
func ( h *handlerCatalog ) Add( bitmask Level, handlers...EntryHandler ) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hmap := h.handlers

	if bitmask&LEVEL_TRACE != 0 {
		hmap[LEVEL_TRACE] = append( hmap[LEVEL_TRACE], handlers... )
	}

	if bitmask&LEVEL_DEBUG != 0 {
		hmap[LEVEL_DEBUG] = append( hmap[LEVEL_DEBUG], handlers ... )
	}

	if bitmask&LEVEL_INFO != 0 {
		hmap[LEVEL_INFO] = append( hmap[LEVEL_INFO], handlers ... )
	}

	if bitmask&LEVEL_WARNING != 0 {
		hmap[LEVEL_WARNING] = append( hmap[LEVEL_WARNING], handlers ... )
	}

	if bitmask&LEVEL_ERROR != 0 {
		hmap[LEVEL_ERROR] = append( hmap[LEVEL_ERROR], handlers ... )
	}

	if bitmask&LEVEL_FATAL != 0 {
		hmap[LEVEL_FATAL] = append( hmap[LEVEL_FATAL], handlers ... )
	}

	if bitmask&LEVEL_PANIC != 0 {
		hmap[LEVEL_PANIC] = append( hmap[LEVEL_PANIC], handlers ... )
	}
}




type IEntryDispatcher interface {
	NewWorker(Level, WorkerPool, func())
	Close()
	Register( Level, ... EntryHandler )
}

type EntryDispatcher struct {
	workers         disposeChannels
	handlersCatalog handlerCatalog
}
func NewEntryDispatcher() *EntryDispatcher {
	return &EntryDispatcher{
		workers: disposeChannels{
			mu: sync.Mutex{},
			channels: make([]chan bool, 0, 100),
		},
		handlersCatalog: handlerCatalog{
			mu: sync.Mutex{},
			handlers: make(map[Level][]EntryHandler),
		},
	}
}

func (d *EntryDispatcher) NewWorker(level Level, pool WorkerPool, done func()) {
	dispose :=  make(chan bool)
	d.workers.Add( dispose )
	go func(dispose chan bool) {
		queue := make(EntryQueue)
		for {
			pool <- queue

			select {
			case entry := <- queue:
				if ( entry.Level >= level ) {
					handlers := d.handlersCatalog.handlers[entry.Level]
					if ( handlers != nil ) {
						for _, h := range handlers {
							h( entry )
						}
					}
				}
				done()
			case <- dispose:
				close(queue)
				return
			}
		}
	}(dispose)
}
func (d *EntryDispatcher) Close() {
	for _, c := range d.workers.channels {
		c <- true
	}
}

func (d *EntryDispatcher) Register( bitmask Level, handlers ... EntryHandler ) {
	d.handlersCatalog.Add(bitmask, handlers...)
}