package ddd

type Aggregate interface {
	Entity
	AddEvent(event Event)
	GetEvents() []Event
	ClearEvents()
}

type AggregateBase struct {
	ID     string
	events []Event
}

func (a AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a AggregateBase) GetEvents() []Event {
	return a.events
}

func (a *AggregateBase) ClearEvents() {
	a.events = make([]Event, 0)
}
