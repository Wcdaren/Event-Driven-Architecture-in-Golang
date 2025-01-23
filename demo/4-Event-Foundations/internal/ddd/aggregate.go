package ddd

// Aggregate represents a DDD aggregate
type Aggregate interface {
	Entity
	AddEvent(event Event)
	GetEvents() []Event
}

// AggregateBase provides a default implementation for Aggregate
type AggregateBase struct {
	ID     string
	events []Event
}

func (a *AggregateBase) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a AggregateBase) GetEvents() []Event {
	return a.events
}
