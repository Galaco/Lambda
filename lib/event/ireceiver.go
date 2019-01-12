package event

// Receivable is the handler for a dispatched event.
type Receivable func(IEvent)
