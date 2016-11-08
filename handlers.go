package meerkats

import "sync"

type catalog struct {
	mu			sync.Mutex
	handlers  	map[Level][]EntryHandler
}
func newHandlerManager() *catalog {
	return &catalog{
		mu:			sync.Mutex{},
		handlers:	make(map[Level][]EntryHandler, 7),
	}
}
func ( h *catalog ) Add( bitmask Level, handlers...EntryHandler ) {
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