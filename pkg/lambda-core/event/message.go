package event

// Message is a generic event message
// Contains the type of event
type Message struct {
	id MessageType
}

// Type Gets type of event
func (message Message) Type() MessageType {
	return message.id
}
