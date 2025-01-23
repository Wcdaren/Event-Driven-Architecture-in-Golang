// internal/ddd/event.go
package ddd

import "context"

// Event represents a domain event
type Event interface {
	EventName() string
}

// EventHandler handles domain events
type EventHandler func(ctx context.Context, event Event) error

