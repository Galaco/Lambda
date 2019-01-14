package event

var singleton *dispatcher

// dispatcher is a very basic event subscribe/dispatch manager
// it exposes a Singleton that anything can subscribe to or publish
// events to.
// Events are processed FIFO, and dispatched to subscribers in subscription
// order, and it runs within its own goroutine.
//
// Since callbacks are handled in the context of a routine, any actions that are required to
// run within a specific threads context must be used cautiously.
//
// The main purpose of the dispatcher is to provide an event management system to decouple
// and allow the views and controllers to safely communicate with each other
// asynchronously.
type dispatcher struct {
	running   bool
	receivers map[string][]Receivable

	messages []IEvent
	//mutex    sync.Mutex
}

// Initialize starts the singleton running in a separate routine.
func (dispatch *dispatcher) Initialize() {
	if dispatch.running == true {
		return
	}
	dispatch.running = true
	//
	//go func() {
	//	for dispatch.running == true {
	//		dispatch.mutex.Lock()
	//		if len(dispatch.messages) == 0 {
	//			dispatch.mutex.Unlock()
	//			time.Sleep(time.Millisecond)
	//			continue
	//		}
	//		dispatch.processMessages()
	//		dispatch.mutex.Unlock()
	//	}
	//}()
}

func (dispatch *dispatcher) processMessages() {
	if len(dispatch.messages) == 0 {
		return
	}

	for i := 0; i < len(dispatch.messages); i++ {
		message := dispatch.messages[0]
		dispatch.messages = dispatch.messages[1:]
		if dispatch.receivers[message.Type()] != nil {
			for _, receiver := range dispatch.receivers[message.Type()] {
				receiver(message)
			}
		}
	}

	//for _, message := range dispatch.messages {
	//	if dispatch.receivers[message.Type()] != nil {
	//		for _, receiver := range dispatch.receivers[message.Type()] {
	//			receiver(message)
	//		}
	//	}
	//}
	//dispatch.messages = make([]IEvent, 0)
}

// Close tells the dispatcher to finish running.
// It will execute any events currently in the queue first.
func (dispatch *dispatcher) Close() {
	dispatch.running = false
}

// Dispatch notified the dispatcher that the specified event has occurred,
// and appends it to the end of the current dispatch queue for processing.
func (dispatch *dispatcher) Dispatch(action IEvent) {
	// @TODO This can actually affect dispatch order, but it seems to prevent locking,
	// fix the cause of locking instead of working around by dispatching in a routine.
	//go func() {
	//	dispatch.mutex.Lock()
	//	dispatch.messages = append(dispatch.messages, action)
	//	dispatch.mutex.Unlock()
	//}()
	dispatch.messages = append(dispatch.messages, action)
	dispatch.processMessages()

}

// Subscribe registers a callback against a particular event.
// Whenever a specified event occurs, the callback is executed.
func (dispatch *dispatcher) Subscribe(action string, receiver Receivable) {
	//dispatch.mutex.Lock()
	if dispatch.receivers[action] == nil {
		dispatch.receivers[action] = make([]Receivable, 0)
	}
	dispatch.receivers[action] = append(dispatch.receivers[action], receiver)
	//dispatch.mutex.Unlock()
}

// Singleton returns the dispatcher.
// There should only be 1 event dispatcher, which should run in the background for the
// lifetime of an application.
func Singleton() *dispatcher {
	if singleton == nil {
		singleton = &dispatcher{
			receivers: map[string][]Receivable{},
		}
	}
	return singleton
}
