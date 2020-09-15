package event

// IMessage Generic event manager message interface
// All messages need to implement this
type IMessage interface {
	Type() MessageType
}
