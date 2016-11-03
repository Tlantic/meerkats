package meerkats

import "sync"

type WorkerPool struct {
	mu    					sync.Mutex
	availableQueues			chan chan Entry
	close 					[]chan bool
}
func NewWorkerPool(initCapacity int) *WorkerPool {
	return &WorkerPool{
		mu:					sync.Mutex{},
		availableQueues: 	make(chan chan Entry),
		close: 				make([]chan bool, 0, initCapacity),
	}
}
func (d *WorkerPool) Add( ch chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.close = append( d.close, ch )
}
func (d *WorkerPool) Close() {
	for _, c := range d.close {
		c <- true
	}
}
