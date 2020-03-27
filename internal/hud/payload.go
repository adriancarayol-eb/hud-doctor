package hud

import "time"

// Payload is a struct to represent the data to be rendered in the hud.
type Payload struct {
	Body      string
	TimeStamp time.Time
}
