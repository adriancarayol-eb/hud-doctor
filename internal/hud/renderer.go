package hud

import (
	tm "github.com/buger/goterm"
	"log"
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
	_, err := tm.Println(tm.Bold("yak doctor"))

	if err != nil {
		log.Printf("Error: %v", err)
	}

	_, err = tm.Println(table)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	tm.Flush()
}
