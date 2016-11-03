package meerkats

import "sync"

type workperPool struct {
	mu         		sync.Mutex
	queueChans 		chan chan Entry
	closeChans 		[]chan bool
}

func newWorkerPool(initCapacity int) *workperPool {
	return &workperPool{
		mu:			sync.Mutex{},
		queueChans:	make(chan chan Entry),
		closeChans:	make([]chan bool, 0, initCapacity),
	}
}
func (d *workperPool) AddCloseChan(ch chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.closeChans = append(d.closeChans, ch)
}
func (d *workperPool) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, c := range d.closeChans {
		c <- true
	}
}
