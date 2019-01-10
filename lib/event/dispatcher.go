package event

import "sync"

var singleton *dispatcher

type dispatcher struct {
	running bool
	receivers map[string][]Receivable

	messages []IAction
	mutex sync.Mutex
}

func (dispatch *dispatcher) Initialize() {
	dispatch.running = true

	go func() {
		for dispatch.running == true {
			dispatch.mutex.Lock()
			for _,message := range dispatch.messages {
				if dispatch.receivers[message.Type()] != nil {
					for _,receiver := range dispatch.receivers[message.Type()] {
						receiver(message)
					}
				}
			}
			dispatch.mutex.Unlock()
		}
	}()
}

func (dispatch *dispatcher) Close() {
	dispatch.running = false
}

func (dispatch *dispatcher) Dispatch(action IAction) {
	dispatch.mutex.Lock()
	dispatch.messages = append(dispatch.messages, action)
	dispatch.mutex.Unlock()
}

func (dispatch *dispatcher) Listen(action string, receiver Receivable) {
	dispatch.mutex.Lock()
	if dispatch.receivers[action] == nil {
		dispatch.receivers[action] = make([]Receivable, 0)
	}
	dispatch.receivers[action] = append(dispatch.receivers[action], receiver)
	dispatch.mutex.Unlock()
}

func Singleton() *dispatcher {
	if singleton == nil {
		singleton = &dispatcher{
			receivers: map[string][]Receivable{},
		}
	}
	return singleton
}