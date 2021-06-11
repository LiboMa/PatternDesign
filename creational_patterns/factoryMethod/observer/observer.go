// Package main serves as an example application that makes use of the observer pattern.
// Playground: https://play.golang.org/p/cr8jEmDmw0
package observer

import (
	"fmt"
)

type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data int64
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event)
	}

	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of Observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}
)

type (
	EventObserver struct {
		Id int
	}

	EventNotifier struct {
		// Using a map with an empty struct allows us to keep the Observers
		// unique while still keeping memory usage relatively low.
		Observers map[Observer]struct{}
	}
)

func (o *EventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.Id, e.Data)
}

func (o *EventNotifier) Register(l Observer) {
	o.Observers[l] = struct{}{}
}

func (o *EventNotifier) Deregister(l Observer) {
	delete(o.Observers, l)
}

func (p *EventNotifier) Notify(e Event) {
	for o := range p.Observers {
		o.OnNotify(e)
	}
}
