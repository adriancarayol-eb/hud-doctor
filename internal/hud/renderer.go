package hud

import (
	"sync"

	tm "github.com/buger/goterm"
)

// Renderer is a struct to render content in the hud.
type Renderer struct {
	mu *sync.RWMutex
}

// New creates a Renderer used to display content in the hud.
func New() *Renderer {
	return &Renderer{
		mu: &sync.RWMutex{},
	}
}

// Refresh the HUD (console) with the new content.
func (r *Renderer) Refresh(payload Payload) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(payload.TimeStamp)
	tm.Println(payload.Body)
	tm.Flush()
}
