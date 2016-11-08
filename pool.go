package meerkats

import "sync"

type workerManager struct {
	mu         sync.Mutex
	pool       chan chan Entry
	closeChans []chan bool
}

func newWorkerManager(initCapacity uint) *workerManager {
	return &workerManager{
		mu:			sync.Mutex{},
		pool:		make(chan chan Entry),
		closeChans:	make([]chan bool, 0, initCapacity),
	}
}
func (d *workerManager) AddCloseChan(ch chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.closeChans = append(d.closeChans, ch)
}
func (d *workerManager) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, c := range d.closeChans {
		c <- true
	}
	close(d.pool)
}
