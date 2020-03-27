package hud

import (
	tm "github.com/buger/goterm"
)

// Renderer is a struct to render content in the hud.
type Renderer struct {

}

// New creates a Renderer used to display content in the hud.
func New() *Renderer {
	return &Renderer{}
}

// Refresh the HUD (console) with the new content.
func (r *Renderer) Refresh(table *tm.Table) {
	tm.Clear() // Clear current screen
	tm.MoveCursor(1, 1)
	tm.Println(tm.Bold("yak doctor"))
	tm.Println(table)
	tm.Flush()
}
