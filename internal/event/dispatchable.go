package event

// Dispatchable represents a dispatchable event.
// An implementing struct would normally include any payload as
// properties.
type Dispatchable interface {
	Type() string
}
