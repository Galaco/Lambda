package event

// MessageType
type MessageType string

// EventHandle
type EventHandle uint

var eventHandleCounter EventHandle

// newEventHandle
func newEventHandle() EventHandle {
	eventHandleCounter++
	return eventHandleCounter
}
