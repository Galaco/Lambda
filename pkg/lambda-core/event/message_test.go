package event

import "testing"

func TestMessage_GetType(t *testing.T) {
	msg := Message{
		id: MessageType("foo"),
	}
	if msg.Type() != MessageType("foo") {
		t.Errorf("Unexpected message type. Got %s, but expected %s", MessageType("foo"), msg.Type())
	}
}
