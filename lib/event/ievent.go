package event

// IEvent represents a dispatchable event.
// An implementing struct would normally include any payload as
// properties.
type IEvent interface {
	Type() string
}
